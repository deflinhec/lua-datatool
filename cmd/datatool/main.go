package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	doc "datatool.deflinhec.dev/doc"
	"datatool.deflinhec.dev/editor"
	gext "datatool.deflinhec.dev/internal/giuext"
	g "github.com/AllenDang/giu"
	"github.com/jessevdk/go-flags"
	"github.com/sqweek/dialog"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "datatool.deflinhec.dev/internal/translations"
)

var (
	Version = "0.0.0"
	Build   = "-"
)

var opts struct {
	File string `long:"launch" short:"l" description:"file to launch"`

	Profiler func() `long:"pprof" short:"p" description:"Enable profiler"`

	Version func() `long:"version" short:"v" description:"檢視建置版號"`
}

var document doc.File
var widget editor.Widget
var translator *message.Printer
var languages = []string{"zh-TW", "zh-CN", "en-GB"}
var parser = flags.NewParser(&opts, flags.Default)

func init() {
	opts.Version = func() {
		fmt.Printf("Version: %v", Version)
		fmt.Printf("\tBuild: %v", Build)
		os.Exit(0)
	}
	opts.Profiler = func() {
		go func() {
			fmt.Println("Profiler listening on port 6060")
			fmt.Println(http.ListenAndServe(":6060", nil))
		}()
	}
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	widget = editor.Dummy()
	onTranslate(languages[0])
}

func onTranslate(lang string) {
	l := language.MustParse(lang)
	translator = message.NewPrinter(l)
	editor.Translator(translator)
	g.Update()
}

func onDrop(names []string) {
	if len(names) == 0 {
		return
	}
	filename := names[0]
	ext := filepath.Ext(filename)
	if !strings.Contains(ext, "lua") {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("副檔名不符")).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	openFile(filename)
	g.Update()
}

func onOpenFile() {
	filename, err := dialog.File().
		Filter("Lua files", "lua").
		Load()
	if err != nil {
		log.Printf("載入失敗: %v", err)
		return
	}
	openFile(filename)
}

func openFile(filename string) {
	abspath, _ := filepath.Abs(filename)
	file, err := doc.Open(abspath)
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("無法載入 %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	value, err := doc.Read(file)
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("無法載入 %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	document = file
	defer widget.Sync()
	switch value := value.(type) {
	case map[string]interface{}:
		widget = editor.Scope().Values(value).Editable()
	default:
		widget = editor.Unmarshal(value)
	}
	log.Println(translator.Sprintf("載入成功 %v", abspath))
}

func onSaveFile() {
	err := doc.Write(document, widget.Interface())
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("儲存失敗: %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	g.Msgbox(translator.Sprintf("訊息"),
		translator.Sprintf("儲存成功")).
		Buttons(g.MsgboxButtonsOk)
}

func onSaveNewFile() {
	filename, err := dialog.File().
		Filter(".lua", "lua").
		Title(translator.Sprintf("另存新檔")).
		SetStartFile(document.Field()).
		Save()
	if err != nil {
		log.Println(translator.Sprintf("儲存失敗: %v", err))
		return
	}
	abspath, _ := filepath.Abs(filename)
	file, err := doc.Open(abspath, doc.WithDocument(document))
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("儲存失敗: %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	err = doc.Write(file, widget.Interface())
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("儲存失敗: %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	document = file
	g.Msgbox(translator.Sprintf("訊息"),
		translator.Sprintf("儲存成功: %v", abspath)).
		Buttons(g.MsgboxButtonsOk)
}

func loop() {
	g.SingleWindowWithMenuBar().Layout(
		g.PrepareMsgbox(),
		g.MenuBar().Layout(
			g.Menu(translator.Sprintf("檔案")).Layout(
				g.MenuItem(translator.Sprintf("開啟")).
					OnClick(onOpenFile),
				g.MenuItem(translator.Sprintf("儲存")).
					Enabled(document != nil).
					OnClick(onSaveFile),
				g.Separator(),
				g.MenuItem(translator.Sprintf("另存新檔")).
					Enabled(document != nil).
					OnClick(onSaveNewFile),
				g.Separator(),
			),
			g.Menu(translator.Sprintf("編輯")).Layout(
				g.MenuItem(translator.Sprintf("開啟")).
					OnClick(onOpenFile),
			).Enabled(document != nil),
			g.Menu(translator.Sprintf("設定")).Layout(
				gext.ComboText(translator.Sprintf("語系"), languages).
					OnSelected(onTranslate),
			),
			g.Align(g.AlignRight).To(
				g.Condition(document != nil, g.Layout{
					g.Custom(func() {
						g.Row(
							g.Label(translator.Sprintf("路徑：%v", document.Path())),
							g.Label(translator.Sprintf("模塊：%v", document.Module())),
							g.Label(translator.Sprintf("欄位：%v", document.Field())),
							g.Label(fmt.Sprintf("md5：%v", document.Md5Sum())),
						).Build()
					}),
				}, g.Layout{}),
			),
		),
		g.Separator(),
		widget,
		g.Separator(),
	)
	if document == nil {
		if len(opts.File) > 0 {
			openFile(opts.File)
		} else {
			onOpenFile()
		}
	}
}

func main() {
	title := "Lua table editor "
	title += fmt.Sprintf("Version:%v Build:%v", Version, Build)
	wnd := g.NewMasterWindow(title, 1024, 960, g.MasterWindowFlagsMaximized)
	wnd.SetDropCallback(onDrop)
	wnd.Run(loop)
}

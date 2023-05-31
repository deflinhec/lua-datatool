// Copyright 2023 Deflinhec
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	g "github.com/AllenDang/giu"
	doc "github.com/deflinhec/lua-datatool/doc"
	"github.com/deflinhec/lua-datatool/editor"
	gext "github.com/deflinhec/lua-datatool/internal/giuext"
	"github.com/jessevdk/go-flags"
	"github.com/ncruces/zenity"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/deflinhec/lua-datatool/internal/translations"
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

var document *doc.File
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
	filename, err := zenity.SelectFile(
		zenity.Title("選擇檔案"),
		zenity.FileFilter{
			Name:     "Lua files",
			Patterns: []string{"lua"},
		},
	)
	if err != nil {
		log.Printf("載入失敗: %v", err)
		return
	}
	openFile(filename)
}

func openFile(filename string) {
	abspath, _ := filepath.Abs(filename)
	file, err := doc.OpenFile(abspath)
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("無法載入 %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
	document = file
	defer widget.Sync()
	switch value := document.Value.(type) {
	case map[string]interface{}:
		widget = editor.Scope().Values(value).Editable()
	default:
		widget = editor.Unmarshal(value)
	}
	log.Println(translator.Sprintf("載入成功 %v", abspath))
}

func onSaveFile() {
	f := doc.File{
		FileInfo: document.FileInfo,
		Value:    widget.Interface(),
	}
	err := f.Save()
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
	filename, err := zenity.SelectFileSave(
		zenity.Title(translator.Sprintf("另存新檔")),
		zenity.FileFilter{
			Name:     "Lua files",
			Patterns: []string{"lua"},
		},
	)
	switch err {
	case nil:
		return
	case zenity.ErrCanceled:
	default:
		log.Println(translator.Sprintf("儲存失敗: %v", err))
	}
	abspath, _ := filepath.Abs(filename)
	err = document.SaveTo(abspath)
	if err != nil {
		g.Msgbox(translator.Sprintf("錯誤"),
			translator.Sprintf("儲存失敗: %v", err)).
			Buttons(g.MsgboxButtonsOk)
		return
	}
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
							g.Label(translator.Sprintf("路徑：%v", document.Filename)),
							g.Label(translator.Sprintf("模塊：%v", document.Namespace)),
							g.Label(translator.Sprintf("欄位：%v", document.Field)),
							g.Label(fmt.Sprintf("md5：%v", document.Md5Sum)),
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

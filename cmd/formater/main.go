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
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/deflinhec/lua-datatool/doc"
	"github.com/deflinhec/lua-datatool/editor"
	"github.com/go-test/deep"
	"github.com/jessevdk/go-flags"
	"github.com/ncruces/zenity"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	Version = "0.0.0"
	Build   = "-"
)

var opts struct {
	Paths []string `long:"path" short:"p" description:"lua path or directory"`

	Ignores []string `long:"ignore" short:"i" description:"ingore specific file"`

	DryRun bool `long:"dry-run" description:"perform dry run"`

	Version func() `long:"version" short:"v" description:"檢視建置版號"`
}

var translator *message.Printer
var languages = []string{"zh-TW", "zh-CN", "en-GB"}
var parser = flags.NewParser(&opts, flags.Default)

func init() {
	opts.Version = func() {
		fmt.Printf("Version: %v", Version)
		fmt.Printf("\tBuild: %v", Build)
		os.Exit(0)
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
	l := language.MustParse(languages[0])
	translator = message.NewPrinter(l)
	editor.Translator(translator)
}

func main() {

	if len(opts.Paths) == 0 {
		files, err := zenity.SelectFileMultiple(
			zenity.Title(translator.Sprintf("格式化檔案")),
			zenity.FileFilter{
				Name:     "Lua files",
				Patterns: []string{"lua"},
			},
		)
		switch err {
		case nil:
			opts.Paths = files
		case zenity.ErrUnsupported:
			log.Println("zenity unsupported")
		case zenity.ErrCanceled:
			os.Exit(0)
		default:
			log.Print(translator.Sprintf("載入失敗: %v", err))
		}
		switch zenity.Question(translator.Sprintf("覆寫原始檔案?"),
			zenity.OKLabel(translator.Sprintf("確認")),
			zenity.CancelLabel(translator.Sprintf("取消"))) {
		case nil:
			opts.DryRun = false
		case zenity.ErrCanceled:
			opts.DryRun = true
		}
	}

	ignores := make(map[string]bool)
	for _, file := range opts.Ignores {
		ignores[strings.TrimSpace(file)] = true
	}

	paths := make([]string, 0)
	for _, path := range opts.Paths {
		if filepath.Ext(path) == ".lua" {
			paths = append(paths, path)
		} else if err := filepath.Walk(path,
			func(path string, info fs.FileInfo, err error) error {
				switch filepath.Ext(path) {
				case ".lua":
					abspath, _ := filepath.Abs(path)
					name := filepath.Base(abspath)
					if _, ok := ignores[name]; ok {
						return nil
					}
					paths = append(paths, abspath)
				}
				return nil
			}); err != nil {
			log.Panic(err)
		}
	}

	wc := sync.WaitGroup{}
	for _, path := range paths {
		wc.Add(1)
		file := path
		go func() {
			defer wc.Done()
			f, err := doc.OpenFile(file)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			if opts.DryRun {
				tmpfile := strings.TrimSuffix(file, ".lua") + ".tmp.lua"
				err = f.SaveTo(tmpfile)
				if err != nil {
					log.Println("[warn]", file, err)
					return
				}
				tf, err := doc.OpenFile(tmpfile)
				if err != nil {
					log.Println("[warn]", file, err)
					return
				}
				log.Println("[done]", filepath.Base(tmpfile))
				if diff := deep.Equal(f.Value, tf.Value); diff != nil {
					log.Println("[warn]", file, "mismatch")
					return
				}
				return
			}
			f, err = doc.OpenFile(file)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			err = f.Save()
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			log.Println("[done]", filepath.Base(file))
		}()
	}
	wc.Wait()
	log.Println("complete files:", len(paths))
}

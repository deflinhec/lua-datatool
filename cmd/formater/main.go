package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"datatool.deflinhec.dev/doc"
	"github.com/jessevdk/go-flags"
)

var (
	Version = "0.0.0"
	Build   = "-"
)

var opts struct {
	Dir string `long:"dir" short:"d" description:"lua directory" default:"."`

	Version func() `long:"version" short:"v" description:"檢視建置版號"`
}

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
}

func metadata(path string, args ...interface{}) error {
	for i, arg := range args {
		b, err := json.MarshalIndent(arg, "", " ")
		if err != nil {
			return err
		}
		abspath, _ := filepath.Abs(path)
		ext := filepath.Ext(abspath)
		name := filepath.Base(path)
		name = strings.ReplaceAll(name, ext, "")
		name = fmt.Sprintf("%v.%v.meta", name, i)
		abspath = filepath.Dir(abspath)
		abspath = filepath.Join(abspath, name)
		f, err := os.Create(abspath)
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(b)
	}
	return nil
}

func main() {
	paths := make([]string, 0)
	filter := func(path string, info fs.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".lua":
			abspath, _ := filepath.Abs(path)
			paths = append(paths, abspath)
		}
		return nil
	}
	wc := sync.WaitGroup{}
	filepath.Walk(opts.Dir, filter)
	abspath, _ := filepath.Abs(opts.Dir)
	log.Println("scanned:", abspath)
	log.Println("scanned files:", len(paths))
	for _, path := range paths {
		wc.Add(1)
		file := path
		go func() {
			defer wc.Done()
			f, err := doc.Open(file)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			a, err := doc.Read(f)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			f, err = doc.Open(file, doc.WithDocument(f))
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			err = doc.Write(f, a)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			f, err = doc.Open(file)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			b, err := doc.Read(f)
			if err != nil {
				log.Println("[warn]", file, err)
				return
			}
			if !reflect.DeepEqual(a, b) {
				if err = metadata(file, a, b); err != nil {
					log.Println("[warn]", file, err)
				}
				log.Println("[warn]", file, "mismatch")
			}
			log.Println("[done]", filepath.Base(file))
		}()
	}
	wc.Wait()
	log.Println("complete files:", len(paths))
}

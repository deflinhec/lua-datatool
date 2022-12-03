package doc

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var translator *message.Printer

func init() {
	lang := language.MustParse("zh-TW")
	translator = message.NewPrinter(lang)
}

func Translator(p *message.Printer) {
	translator = p
}

type File interface {
	Path() string

	Module() string

	Field() string

	Md5Sum() string

	read() (interface{}, error)

	write(value interface{}) error
}

func md5Sum(value interface{}) string {
	hash := md5.Sum([]byte(fmt.Sprint(value)))
	return hex.EncodeToString(hash[:])
}

func Open(path string, opts ...Option) (File, error) {
	abspath, _ := filepath.Abs(path)
	info := fileInfo{
		dir:  filepath.Dir(abspath),
		name: filepath.Base(abspath),
	}
	for _, opt := range opts {
		opt.apply(&info)
	}
	switch filepath.Ext(abspath) {
	case ".lua":
		return &luaFile{fileInfo: info}, nil
	}
	return nil, errors.New(translator.Sprintf("不支援 %v",
		filepath.Ext(abspath)))
}

func Write(file File, value interface{}) error {
	return file.write(value)
}

func Read(file File) (interface{}, error) {
	return file.read()
}

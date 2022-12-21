package doc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
)

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
	return nil, fmt.Errorf("unsupport %v", filepath.Ext(abspath))
}

func Write(file File, value interface{}) error {
	return file.write(value)
}

func Read(file File) (interface{}, error) {
	return file.read()
}

package assets

import (
	"bytes"
	"embed"
	"image"
	"image/png"
)

var (
	//go:embed *.png
	res embed.FS
)

func Image(name string) (image.Image, error) {
	b, err := res.ReadFile(name)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return png.Decode(r)
}

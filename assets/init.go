package assets

import (
	"bytes"
	"image"
	"image/png"
)

func Image(name string) (image.Image, error) {
	b, err := Asset(name)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return png.Decode(r)
}

package editor

import (
	"image"
	"log"

	"datatool.deflinhec.dev/assets"

	g "github.com/AllenDang/giu"
)

func Image(name string) *image.RGBA {
	img, err := assets.Image(name)
	if err != nil {
		log.Panic(err)
	}
	return g.ImageToRgba(img)
}

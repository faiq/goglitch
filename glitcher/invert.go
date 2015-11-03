package glitcher

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func InvertImage(input string, outfile string) {
	reader, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	inverted := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r = 255 - r>>8
			g = 255 - g>>8
			b = 255 - b>>8
			inverted.SetNRGBA(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	pwd, _ := os.Getwd()
	fileName := path.Join(pwd, outfile)
	fp, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	idx := strings.LastIndex(fileName, ".")
	if idx == -1 {
		panic(errors.New("you dont have a supported output file"))
	}
	ext := fileName[idx+1:]
	if ext == "png" {
		err = png.Encode(fp, inverted)
		if err != nil {
			panic(err)
		}
	} else if ext == "jpeg" || ext == "jpg" {
		var opt jpeg.Options
		opt.Quality = 100
		err = jpeg.Encode(fp, inverted, &opt)
		if err != nil {
			panic(err)
		}
	} else if ext == "gif" {
		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(fp, inverted, &opt)
	} else {
		panic(errors.New("you dont have a supported output file"))
	}
}

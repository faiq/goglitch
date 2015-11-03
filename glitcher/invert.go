package glitcher

import (
	"errors"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func InvertImage(input *os.File, outfile string) {
	reader, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, _, err := image.Read(reader)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	inverted := NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			r = r - 255
			g = g - 255
			b = b - 255
			inverted.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	pwd, _ := os.Getwd()
	fileName, _ := path.Join(pwd, outfile)
	fp, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	ext := strings.LastIndex(fp.Name(), ".")
	if ext == "png" {
		err = png.Encode(fp, inverted)
		if err != nil {
			panic(err)
		}
	} else if ext == "jpeg" {
		var opt jpeg.Options
		opt.Quality = 80
		err = jpeg.Encode(fp, inverted, &opt) // put quality to 80%
		if err != nil {
			panic(err)
		}
	} else if ext == "gif" {
		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(out, inverted, &opt)
	} else {
		panic(errors.New("you dont have a supported output file"))
	}
}

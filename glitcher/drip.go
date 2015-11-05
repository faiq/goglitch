package glitcher

import (
	"errors"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"sort"
	"strings"
)

func DripImage(input string, outfile string) {
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
	dripped := image.NewNRGBA(bounds)
	/* at a certain point you want the image to retain
	its color values then as you go lower down you
	want it to start trickling into pixels beneath
	it */
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		var vals []int
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			val := uint32(0 | uint32(r)<<8 | uint32(g)<<16 | uint32(b)<<24)
			vals = append(vals, int(val))
		}
		sort.Ints(vals)
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r := vals[x] >> 8
			g := vals[x] >> 16
			b := vals[x] >> 24
			dripped.SetNRGBA(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(255)})
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
		err = png.Encode(fp, dripped)
		if err != nil {
			panic(err)
		}
	} else if ext == "jpeg" || ext == "jpg" {
		var opt jpeg.Options
		opt.Quality = 100
		err = jpeg.Encode(fp, dripped, &opt)
		if err != nil {
			panic(err)
		}
	} else if ext == "gif" {
		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(fp, dripped, &opt)
	} else {
		panic(errors.New("you dont have a supported output file"))
	}
}

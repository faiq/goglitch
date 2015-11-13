package glitcher

import (
	"github.com/faiq/goglitch/helpers"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sort"
)

func VerticalSort(input string, outfile string) error {
	reader, err := os.Open(input)
	if err != nil {
		return err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bounds := img.Bounds()
	sortedImg := image.NewNRGBA(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var vals []int
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			val := uint32(0 | uint32(r)<<8 | uint32(g)<<16 | uint32(b)<<24)
			vals = append(vals, int(val))
		}
		sort.Ints(vals)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r := vals[y] >> 8
			g := vals[y] >> 16
			b := vals[y] >> 24
			sortedImg.SetNRGBA(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(255)})
		}
	}
	err = helpers.SaveImage(sortedImg, outfile)
	if err != nil {
		return err
	}
	return nil
}

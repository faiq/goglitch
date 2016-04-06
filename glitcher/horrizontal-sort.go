package glitcher

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"sort"

	"github.com/faiq/goglitch/helpers"
)

func HorizontalSort(reader io.Reader, dest io.Writer) error {
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	bounds := img.Bounds()
	dripped := image.NewNRGBA(bounds)
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
	err = helpers.WriteImage(dripped, "png", dest)
	if err != nil {
		return err
	}
	return nil
}

package glitcher

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

type PixelGradient struct {
	Degrees   float64 //representing the direction its going in
	Magnitude float64 //representing the
}

/* every pixel in the picture is going to have 2 values associated with it.
1. gradient intensity
2. gradient direction
to represent this we could use a hash that maps
y*x (int) -> GradientType
*/
func getLumonosity(r, g, b, a uint32) float64 {
	return 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b) + 0*a
}

func computeImageGradient(img image.Image) map[int]PixelGradient {
	ret := make(map[int]PixelGradient)
	// we're going to choose to ignore the border pixels right now
	// rather than applying a padding or a stretch
	bounds := img.Bounds()
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			pix := y*bounds.Max.Y + x
			northLum := getLumonosity(img.At(x, y-1).RGBA())
			southLum := getLumonosity(img.At(x, y+1).RGBA())
			eastLum := getLumonosity(img.At(x+1, y).RGBA())
			westLum := getLumonosity(img.At(x-1, y).RGBA())
			dx := eastLum - westLum
			dy := northLum - southLum
			ret[pix] = PixelGradient{
				Degrees:   math.Atan2(dy, dx) * 180 / math.Pi, //Radians -> degrees
				Magnitude: math.Hypot(dy, dx),
			}
		}
	}
	return ret
}

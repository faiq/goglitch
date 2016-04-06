package helpers

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func WriteImage(img image.Image, ext string, w io.Writer) error {
	var err error
	if ext == "png" {
		err = png.Encode(w, img)
		if err != nil {
			return err
		}
	} else if ext == "jpeg" || ext == "jpg" {
		var opt jpeg.Options
		opt.Quality = 100
		err = jpeg.Encode(w, img, &opt)
		if err != nil {
			return err
		}
	} else if ext == "gif" {
		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(w, img, &opt)
	} else {
		return errors.New("you dont have a supported output file")
	}
	return nil
}

package helpers

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func SaveImage(img image.Image, outfile string) error {
	pwd, _ := os.Getwd()
	fileName := path.Join(pwd, outfile)
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	idx := strings.LastIndex(fileName, ".")
	if idx == -1 {
		return errors.New("you dont have a supported output file")
	}
	ext := fileName[idx+1:]
	if ext == "png" {
		err = png.Encode(fp, img)
		if err != nil {
			return err
		}
	} else if ext == "jpeg" || ext == "jpg" {
		var opt jpeg.Options
		opt.Quality = 100
		err = jpeg.Encode(fp, img, &opt)
		if err != nil {
			return err
		}
	} else if ext == "gif" {
		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(fp, img, &opt)
	} else {
		return errors.New("you dont have a supported output file")
	}
	return nil
}

package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func ReadImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	im, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}

	return im, nil
}

func WritePng(im image.Image, path string) error {
	wr, _ := os.Create("result.png")
	err := png.Encode(wr, im)
	if err != nil {
		return err
	}
	return nil
}

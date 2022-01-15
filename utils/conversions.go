package utils

import (
	"image"
	"image/color"
)

func AsGrayImage(im image.Image) *image.Gray {
	bounds := im.Bounds()
	gim := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gim.Set(x, y, im.At(x, y))
		}
	}
	return gim
}

func Threshold(gim *image.Gray, t uint8) *image.Gray {

	bounds := gim.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if gim.GrayAt(x, y).Y < t {
				gim.Set(x, y, color.Black)
			} else {
				gim.Set(x, y, color.White)
			}
		}
	}

	return gim
}

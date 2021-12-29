package utils

import "image"

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

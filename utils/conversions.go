package utils

import (
	d "github.com/vlad-a-barbu/gocr/drawing"
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

func DrawingToImage(drawing d.Drawing) *image.Gray {

	upLeft := image.Point{}
	lowRight := image.Point{X: drawing.Width, Y: drawing.Height}

	gim := image.NewGray(image.Rectangle{Min: upLeft, Max: lowRight})
	bounds := gim.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gim.Set(x, y, color.Gray{Y: 255})
		}
	}

	points := make(map[image.Point]bool)
	for _, line := range drawing.Lines {
		for _, point := range line.Points {
			X := int(point.X)
			Y := int(point.Y)
			gim.Set(X, Y, color.Gray{Y: 0})
			points[image.Point{X, Y}] = true
		}
	}

	Interpolate(gim, points)

	return gim
}

func Interpolate(gim *image.Gray, points map[image.Point]bool) {
	bounds := gim.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := image.Point{X: x, Y: y}
			if gim.GrayAt(x, y).Y == 255 &&
				HasNeighboursFilter(p, gim, points) {
				gim.Set(x, y, color.Gray{Y: 0})
			}
		}
	}
}

func HasNeighboursFilter(p image.Point, gim *image.Gray, points map[image.Point]bool) bool {
	return IsBlack(image.Point{X: p.X + 1, Y: p.Y + 1}, gim, points) ||
		IsBlack(image.Point{X: p.X + 1, Y: p.Y - 1}, gim, points) ||
		IsBlack(image.Point{X: p.X - 1, Y: p.Y - 1}, gim, points) ||
		IsBlack(image.Point{X: p.X - 1, Y: p.Y + 1}, gim, points) ||
		IsBlack(image.Point{X: p.X, Y: p.Y - 1}, gim, points) ||
		IsBlack(image.Point{X: p.X, Y: p.Y + 1}, gim, points) ||
		IsBlack(image.Point{X: p.X + 1, Y: p.Y}, gim, points) ||
		IsBlack(image.Point{X: p.X - 1, Y: p.Y}, gim, points)
}

func IsBlack(p image.Point, gim *image.Gray, points map[image.Point]bool) bool {
	bounds := gim.Bounds()
	_, ok := points[p]
	return ok &&
		p.X < bounds.Max.X &&
		p.X >= bounds.Min.X &&
		p.Y >= bounds.Min.Y &&
		p.Y < bounds.Max.Y &&
		gim.GrayAt(p.X, p.Y).Y == 0
}

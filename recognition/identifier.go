package recognition

import (
	"image"
	"image/color"
)

func MinMax(ps []image.Point, w int, h int) (min image.Point, max image.Point) {

	minx := w - 1
	maxx := 0
	miny := h - 1
	maxy := 0

	for _, p := range ps {
		if p.X > maxx {
			maxx = p.X
		} else if p.X < minx {
			minx = p.X
		}

		if p.Y > maxy {
			maxy = p.Y
		} else if p.Y < miny {
			miny = p.Y
		}
	}

	return image.Point{minx, miny}, image.Point{maxx, maxy}
}

func SubImage(ps []image.Point, gim *image.Gray) image.Image {
	bounds := gim.Bounds()
	h := bounds.Max.Y - bounds.Min.Y
	w := bounds.Max.X - bounds.Min.X
	min, max := MinMax(ps, w, h)
	r := image.Rectangle{min, max}
	return gim.SubImage(r)
}

func TraverseRows(im image.Image) [][]int {
	vals := [][]int{}
	bounds := im.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		v := []int{}
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			p := im.At(x, y)
			Y := color.GrayModel.Convert(p).(color.Gray).Y
			if Y < MAX_Y {
				v = append(v, 1)
			} else {
				v = append(v, 0)
			}
		}
		vals = append(vals, v)
	}
	return vals
}

func TraverseCols(im image.Image) [][]int {
	vals := [][]int{}
	bounds := im.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		v := []int{}
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := im.At(x, y)
			Y := color.GrayModel.Convert(p).(color.Gray).Y
			if Y < MAX_Y {
				v = append(v, 1)
			} else {
				v = append(v, 0)
			}
		}
		vals = append(vals, v)
	}
	return vals
}

func Identify(ps []image.Point) rune {
	return '0'
}

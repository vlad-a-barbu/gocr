package recognition

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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
	si := gim.SubImage(r)
	return si
}

func TraverseRows(im image.Image) ([][]int, []int) {
	cnt := []int{}
	vals := [][]int{}
	bounds := im.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		c := 0
		v := []int{}
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			p := im.At(x, y)
			Y := color.GrayModel.Convert(p).(color.Gray).Y
			if Y < MAX_Y {
				v = append(v, 1)
				c++
			} else {
				v = append(v, 0)
			}
		}
		cnt = append(cnt, c)
		vals = append(vals, v)
	}
	return vals, cnt
}

func TraverseCols(im image.Image) ([][]int, []int) {
	cnt := []int{}
	vals := [][]int{}
	bounds := im.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		c := 0
		v := []int{}
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := im.At(x, y)
			Y := color.GrayModel.Convert(p).(color.Gray).Y
			if Y < MAX_Y {
				v = append(v, 1)
				c++
			} else {
				v = append(v, 0)
			}
		}
		cnt = append(cnt, c)
		vals = append(vals, v)
	}
	return vals, cnt
}

func Histogram(xs []int, name string, id int) {

	xyzs := make(plotter.XYZs, len(xs))
	for i, x := range xs {
		xyzs = append(xyzs, plotter.XYZ{X: float64(i), Y: float64(x), Z: 1})
	}

	p := plot.New()
	p.Title.Text = name
	h1, err := plotter.NewHistogram(xyzs, len(xs))
	if err != nil {
		log.Panic(err)
	}
	p.Add(h1)

	err = p.Save(200, 200, "hists/"+fmt.Sprint(id)+"/"+name+".png")
	if err != nil {
		log.Panic(err)
	}

}

func GenerateHists(im image.Image, id int) {
	_, rc := TraverseCols(im)
	Histogram(rc, "rows", id)

	_, cc := TraverseRows(im)
	Histogram(cc, "cols", id)
}

func Identify(im image.Image) rune {
	return '0'
}

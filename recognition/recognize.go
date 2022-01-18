package recognition

import (
	"github.com/vlad-a-barbu/gocr/models"
	"image"
)

const MAX_Y = 200

func Candidate(p image.Point, gim *image.Gray, lookup map[image.Point]int) bool {
	bounds := gim.Bounds()
	_, exists := lookup[p]
	return p.X < bounds.Max.X &&
		p.X >= bounds.Min.X &&
		p.Y >= bounds.Min.Y &&
		p.Y < bounds.Max.Y &&
		gim.GrayAt(p.X, p.Y).Y < MAX_Y &&
		!exists
}

func SubElement(gim *image.Gray, lookup map[image.Point]int, p image.Point, eid *int) *[]image.Point {
	ps := []image.Point{}
	Fill(gim, lookup, p, *eid, &ps)
	if len(ps) > 0 {
		*eid++
		return &ps
	}
	return nil
}

func Fill(gim *image.Gray, lookup map[image.Point]int, p image.Point, eid int, ps *[]image.Point) {
	if Candidate(p, gim, lookup) {
		lookup[p] = eid
		*ps = append(*ps, p)
		Fill(gim, lookup, image.Point{p.X + 1, p.Y}, eid, ps)
		Fill(gim, lookup, image.Point{p.X - 1, p.Y}, eid, ps)
		Fill(gim, lookup, image.Point{p.X, p.Y + 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X, p.Y - 1}, eid, ps)
	} else {
		return
	}
}

func SubElement2(gim *image.Gray, lookup map[image.Point]int, p image.Point, eid *int) *[]image.Point {
	ps := []image.Point{}
	Fill2(gim, lookup, p, *eid, &ps)
	if len(ps) > 0 {
		*eid++
		return &ps
	}
	return nil
}

func Fill2(gim *image.Gray, lookup map[image.Point]int, p image.Point, eid int, ps *[]image.Point) {
	if Candidate(p, gim, lookup) {
		lookup[p] = eid
		*ps = append(*ps, p)
		Fill(gim, lookup, image.Point{p.X + 1, p.Y}, eid, ps)
		Fill(gim, lookup, image.Point{p.X - 1, p.Y}, eid, ps)
		Fill(gim, lookup, image.Point{p.X, p.Y + 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X, p.Y - 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X + 1, p.Y + 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X - 1, p.Y + 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X + 1, p.Y - 1}, eid, ps)
		Fill(gim, lookup, image.Point{p.X - 1, p.Y - 1}, eid, ps)
	} else {
		return
	}
}

func Recognize(gim *image.Gray) map[int][]image.Point {

	eid := -1
	lookup := map[image.Point]int{}
	elems := map[int][]image.Point{}
	bounds := gim.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := image.Point{x, y}
			ps := SubElement(gim, lookup, p, &eid)
			if ps != nil && len(*ps) > 4 {
				elems[eid] = *ps
			}
		}
	}

	return elems
}

type HistData struct {
	Lines   int
	Weights []int
}

func GetHistData(freq_matrix [][]int) []int {

	var data []int
	for _, freqs := range freq_matrix {

		lines := 0
		ok := true
		for _, freq := range freqs {
			if freq == 1 && ok {
				lines++
				ok = false
			}
			if freq == 0 {
				ok = true
			}
		}

		data = append(data, lines)
	}

	return data
}

func Hamming(c1 []int, c2 []int) int {

	var min int
	if len(c1) < len(c2) {
		min = len(c1)
	} else {
		min = len(c2)
	}

	dist := 0

	for i := 0; i < min; i++ {
		if c1[i] != c2[i] {
			dist++
		}
	}

	return dist
}

func Guess(gim *image.Gray) (image.Image, []rune) {

	m := Recognize(gim)
	points := m[0]
	si := SubImage(points, gim)

	rows, _ := TraverseCols(si)
	rdata := GetHistData(rows)

	cols, _ := TraverseRows(si)
	cdata := GetHistData(cols)

	matches := models.Match(rdata, cdata)

	return si, matches

}

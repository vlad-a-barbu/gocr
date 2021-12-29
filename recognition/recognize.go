package recognition

import (
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

func Recognize(gim *image.Gray) map[int][]image.Point {

	eid := 1
	lookup := map[image.Point]int{}
	elems := map[int][]image.Point{}
	bounds := gim.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			p := image.Point{x, y}
			ps := SubElement(gim, lookup, p, &eid)
			if ps != nil && len(*ps) > 0 {
				elems[eid] = *ps
			}
		}
	}

	return elems
}

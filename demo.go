package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"

	r "github.com/vlad-a-barbu/gocr/recognition"
	u "github.com/vlad-a-barbu/gocr/utils"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Image path not provided")
	}
	path := os.Args[1]

	im, err := u.ReadImage(path)
	if err != nil {
		log.Fatal(err)
	}

	gim := u.AsGrayImage(im)

	m := r.Recognize(gim)
	eid, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	bounds := gim.Bounds()
	res := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			res.Set(x, y, gim.GrayAt(x, y))
		}
	}

	for id, ps := range m {
		if id < eid {
			for _, p := range ps {
				res.Set(p.X, p.Y, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	si := r.SubImage(m[eid], gim)
	sb := si.Bounds()
	w := sb.Max.X - sb.Min.X
	h := sb.Max.Y - sb.Min.Y
	char := r.Identify(m[eid])
	fmt.Println("Width: ", w, "; Height: ", h)

	fmt.Println("Cols: ")
	cols := r.TraverseRows(si)
	for _, c := range cols {
		fmt.Println(c)
	}

	fmt.Println("Rows: ")
	rows := r.TraverseCols(si)
	for _, r := range rows {
		fmt.Println(r)
	}

	fmt.Printf("%c\n", char)

	err = u.WritePng(si, "result.png")
	if err != nil {
		log.Fatal(err)
	}
}

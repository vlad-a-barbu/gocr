package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"

	models "github.com/vlad-a-barbu/gocr/models"
	r "github.com/vlad-a-barbu/gocr/recognition"
	server "github.com/vlad-a-barbu/gocr/server"
	u "github.com/vlad-a-barbu/gocr/utils"
)

func main() {
	server.Serve(8081)
}

func Demo(writeHists bool, threshold bool) {
	if len(os.Args) < 3 {
		log.Fatalln("Provide the image path and a subelement number")
	}
	path := os.Args[1]
	sen, _ := strconv.Atoi(os.Args[2])

	im, err := u.ReadImage(path)
	if err != nil {
		log.Fatal(err)
	}

	gim := u.AsGrayImage(im)
	u.WritePng(gim, "grayscale.png")

	if threshold {
		gim := u.Threshold(gim, r.MAX_Y)
		u.WritePng(gim, "blackwhite.png")
	}

	m := r.Recognize(gim)

	r, c := ToHistData(sen, m, gim)
	fmt.Println("Rows: ", r)
	fmt.Println("Cols: ", c)

	res1 := models.Match(r, c)
	for _, r := range res1 {
		fmt.Printf("Match: '%c'\n", r)
	}

	if writeHists {
		u.WriteHists(m, gim)
	}
}

func ToHistData(id int, m map[int][]image.Point, gim *image.Gray) (rd []int, cd []int) {
	points := m[id]
	si := r.SubImage(points, gim)
	u.WritePng(si, "subimage.png")

	rows, _ := r.TraverseCols(si)
	rdata := r.GetHistData(rows)

	cols, _ := r.TraverseRows(si)
	cdata := r.GetHistData(cols)

	return rdata, cdata
}

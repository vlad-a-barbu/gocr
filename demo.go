package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"

	models "github.com/vlad-a-barbu/gocr/models"
	r "github.com/vlad-a-barbu/gocr/recognition"
	u "github.com/vlad-a-barbu/gocr/utils"
)

func main() {

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

	/*gim := u.Threshold(gray, r.MAX_Y)
	u.WritePng(gim, "blackwhite.png")*/

	m := r.Recognize(gim)

	r, c := Test(sen, m, gim)
	fmt.Println("Rows: ", r)
	fmt.Println("Cols: ", c)

	res1 := models.Match(r, c)
	for _, r := range res1 {
		fmt.Printf("Match: '%c'\n", r)
	}

	//u.WriteHists(m, gim)
}

func Test(id int, m map[int][]image.Point, gim *image.Gray) (rd []int, cd []int) {
	points := m[id]
	si := r.SubImage(points, gim)
	u.WritePng(si, "subimage.png")

	rows, _ := r.TraverseCols(si)
	rdata := r.GetHistData(rows)

	cols, _ := r.TraverseRows(si)
	cdata := r.GetHistData(cols)

	return rdata, cdata
}

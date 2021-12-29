package main

import (
	"fmt"
	"log"
	"os"

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

	os.RemoveAll("hists")

	if err := os.Mkdir("hists", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for id, ps := range m {
		if err := os.Mkdir("hists/"+fmt.Sprint(id), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		si := r.SubImage(ps, gim)
		if si.Bounds().Max.X == 0 || si.Bounds().Max.Y == 0 {
			continue
		}

		err = u.WritePng(si, "hists/"+fmt.Sprint(id)+"/image.png")
		if err != nil {
			log.Fatal(err)
		}

		r.GenerateHists(si, id)
		fmt.Println("Hist generated ", id)
	}
}

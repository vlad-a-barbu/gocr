package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(len(m))
	for id, ps := range m {
		go func(i int, points []image.Point) {
			if err := os.Mkdir("hists/"+fmt.Sprint(i), os.ModePerm); err != nil {
				log.Fatal(err)
			}

			si := r.SubImage(points, gim)
			bounds := si.Bounds()
			if bounds.Max.X == 0 || bounds.Max.Y == 0 {
				wg.Done()
				return
			}

			err = u.WritePng(si, "hists/"+fmt.Sprint(i)+"/image.png")
			if err != nil {
				log.Fatal(err)
			}

			r.GenerateHists(si, i)
			fmt.Println("Hist generated ", i)
			wg.Done()
		}(id, ps)
	}
	wg.Wait()
}

package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"sync"

	r "github.com/vlad-a-barbu/gocr/recognition"
)

func ReadImage(path string) (image.Image, error) {

	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(reader *os.File) {
		err := reader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(reader)

	im, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}

	return im, nil
}

func WritePng(im image.Image, path string) error {
	wr, _ := os.Create(path)
	err := png.Encode(wr, im)
	if err != nil {
		return err
	}
	return nil
}

func WriteHists(m map[int][]image.Point, gim *image.Gray) {

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

			err := WritePng(si, "hists/"+fmt.Sprint(i)+"/image.png")
			if err != nil {
				log.Fatal(err)
			}

			r.GenerateHists(si, i)
			fmt.Println("Hist generated ", i)

			f, err := os.Create("hists/" + fmt.Sprint(i) + "/label.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err2 := f.WriteString("UNKNOWN\n")
			if err2 != nil {
				log.Fatal(err2)
			}

			wg.Done()
		}(id, ps)
	}
	wg.Wait()
}

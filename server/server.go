package server

import (
	"encoding/json"
	"fmt"
	d "github.com/vlad-a-barbu/gocr/drawing"
	recog "github.com/vlad-a-barbu/gocr/recognition"
	"github.com/vlad-a-barbu/gocr/utils"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	var drawing d.Drawing

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&drawing)

	if err != nil {
		panic(err)
	}

	err = decoder.Decode(&drawing)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	gim := utils.DrawingToImage(drawing)

	err = utils.WritePng(gim, "drawing.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Drawing converted to PNG and written to disk")
	}

	si, matches := recog.Guess(gim)
	utils.WritePng(si, "guess.png")

	for _, r := range matches {
		fmt.Printf("%c ", r)
	}

	json.NewEncoder(w).Encode(d.Recognition{Matches: matches})
}

func Serve(port int) {
	http.HandleFunc("/recognize", handler)
	fmt.Println("Listening on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

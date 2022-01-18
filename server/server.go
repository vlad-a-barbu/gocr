package server

import (
	"encoding/json"
	"fmt"
	recog "github.com/vlad-a-barbu/gocr/recognition"
	"github.com/vlad-a-barbu/gocr/utils"
	vm "github.com/vlad-a-barbu/gocr/viewmodels"
	"log"
	"net/http"
)

func recognizeHandler(w http.ResponseWriter, r *http.Request) {

	var drawing vm.Drawing

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&drawing)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	gim := utils.DrawingToImage(drawing)

	err = utils.WritePng(gim, "viewmodels.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Drawing converted to PNG and written to disk")
	}

	matches := recog.Guess(gim)

	for _, r := range matches {
		fmt.Printf("%c ", r)
	}

	json.NewEncoder(w).Encode(vm.Recognition{Matches: matches})
}

func matchExprHandler(w http.ResponseWriter, r *http.Request) {

	var expressions vm.Expressions
	var data vm.Drawing

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&expressions)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(expressions.Data), &data)
	if err != nil {
		panic(err)
	}

	gim := utils.DrawingToImage(data)
	err = utils.WritePng(gim, "viewmodels.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Drawing converted to PNG and written to disk")
	}
	result := recog.MatchExpressions(gim, expressions)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(result)
}

func Serve(port int) {
	http.HandleFunc("/recognize", recognizeHandler)
	http.HandleFunc("/matchExpr", matchExprHandler)
	fmt.Println("Listening on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

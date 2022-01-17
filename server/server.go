package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Drawing struct {
	Data string
}

func handler(w http.ResponseWriter, r *http.Request) {

	var drawing Drawing

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&drawing)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data: ", drawing.Data)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(drawing)
}

func Serve(port int) {
	http.HandleFunc("/recognize", handler)
	fmt.Println("Listening on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

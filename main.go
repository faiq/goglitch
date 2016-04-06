package main

import (
	"log"
	"net/http"

	"github.com/faiq/goglitch/glitcher"
)

func hsort(w http.ResponseWriter, req *http.Request) {
	in, _, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer in.Close()
	err = glitcher.HorizontalSort(in, w)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/horizontalsort", hsort)
	mux.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
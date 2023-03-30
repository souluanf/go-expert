package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

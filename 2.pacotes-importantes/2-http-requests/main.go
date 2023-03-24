package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))
	request.Body.Close()
}

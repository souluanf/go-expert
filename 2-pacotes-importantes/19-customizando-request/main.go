package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	requestBody := bytes.NewBuffer([]byte(`{"nome":"Luan"}`))
	req, err := http.NewRequest(http.MethodPost, "https://www.google.com", requestBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

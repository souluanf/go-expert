package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- "Hello, " + nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Luan", hello)
	ler(hello)
}

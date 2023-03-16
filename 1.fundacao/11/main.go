package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	luan := Cliente{
		Nome:  "Luan",
		Idade: 28,
		Ativo: false,
	}
	fmt.Println(luan)
}

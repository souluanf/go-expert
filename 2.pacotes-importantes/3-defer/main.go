package main

import "fmt"

func main() {
	fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha") //Segura a execuçãom e executa no final
	fmt.Println("Terceira Linha")
}

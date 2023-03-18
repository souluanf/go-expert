package main

import (
	"fmt"
	"go-expert/1.fundacao/19/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", soma)
	fmt.Println(matematica.A)
}

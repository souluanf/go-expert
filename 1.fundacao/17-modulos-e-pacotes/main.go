package main

import (
	"fmt"
	"go-expert/1.fundacao/17-modulos-e-pacotes/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", soma)
	fmt.Println(matematica.A)
}

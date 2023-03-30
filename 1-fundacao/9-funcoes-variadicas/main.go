package main

import (
	"fmt"
)

func main() {
	valor := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Valor: %d\n", valor)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

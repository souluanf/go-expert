package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		fmt.Printf("Erro: %s\n", err.Error())
		return
	}
	fmt.Printf("Valor: %d\n", valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("valor acima do permitido")
	}
	return a + b, nil
}

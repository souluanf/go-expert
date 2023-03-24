package main

import "fmt"

func main() {
	salarios := map[string]int{"Luan": 2200, "Joao": 3000, "Maria": 4000, "Jose": 5000}
	fmt.Printf("Salario de Luan: %d\n", salarios["Luan"])
	delete(salarios, "Luan")
	fmt.Printf("Salario de Luan: %d\n", salarios["Luan"])
	salarios["Luan"] = 3000.00
	fmt.Printf("Salario de Luan: %d\n", salarios["Luan"])

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s, Salario: %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}

	//sal := make(map[string]int)
	//sal1 := map[string]int{}

}

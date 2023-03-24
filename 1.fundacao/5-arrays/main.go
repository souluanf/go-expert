package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10
	fmt.Println(meuArray)
	fmt.Println(meuArray[len(meuArray)-1])

	for index, value := range meuArray {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

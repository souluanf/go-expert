package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	tamanho, err := f.Write([]byte("Hello World!"))
	//tamanho, err := f.WriteString("Hello World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	err = f.Close()
	if err != nil {
		return
	}

	//
	// Lendo o arquivo
	//
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	println("Conteúdo do arquivo: ")
	println(string(arquivo))

	//
	// Lendo o arquivo linha a linha
	//
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 12)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s desativado\n", c.Nome)
}

func main() {
	luan := Cliente{Nome: "Luan", Idade: 28, Ativo: true, Endereco: Endereco{Logradouro: "Rua Teste", Numero: 45, Cidade: "São Paulo", Estado: "SP"}}
	luan.Desativar()
}

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (e Empresa) Desativar() {
	fmt.Printf("A empresa %s desativada\n", e.Nome)
}
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	luan := Cliente{Nome: "Luan", Idade: 28, Ativo: true, Endereco: Endereco{Logradouro: "Rua Teste", Numero: 45, Cidade: "São Paulo", Estado: "SP"}}
	minhaEmpresa := Empresa{Nome: "Minha Empresa"}
	Desativacao(luan)
	Desativacao(minhaEmpresa)
}

package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	//t := template.Must(template.ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 120},
		{"Java", 80},
		{"Python", 40},
	})
	if err != nil {
		panic(err)
	}
}

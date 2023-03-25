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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 120},
		{"Java", 80},
		{"Python", 40},
	})
	if err != nil {
		panic(err)
	}
}

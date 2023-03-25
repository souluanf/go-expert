package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 120},
			{"Java", 80},
			{"Python", 40},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe("localhost:8080", nil)
}

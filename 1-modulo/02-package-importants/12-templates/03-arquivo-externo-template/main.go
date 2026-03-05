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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.New("content.html").ParseFiles(templates...))

		curso := Cursos{
			{Nome: "Go Lang", CargaHoraria: 40},
			{Nome: "Python", CargaHoraria: 30},
			{Nome: "JavaScript", CargaHoraria: 20},
		}

		err := temp.Execute(w, curso)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}

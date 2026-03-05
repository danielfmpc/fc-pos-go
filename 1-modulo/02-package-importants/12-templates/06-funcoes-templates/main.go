package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	curso := []Curso{
		{Nome: "Go Lang", CargaHoraria: 40},
		{Nome: "Python", CargaHoraria: 30},
		{Nome: "JavaScript", CargaHoraria: 20},
	}
	temp := template.New("CursoTemplate")

	temp.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})

	temp = template.Must(temp.ParseFiles("template.html"))

	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

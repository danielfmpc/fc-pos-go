package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := []Curso{
		{Nome: "Go Lang", CargaHoraria: 40},
		{Nome: "Python", CargaHoraria: 30},
		{Nome: "JavaScript", CargaHoraria: 20},
	}
	temp := template.Must(template.New("CursoTemplate").ParseFiles("template.html"))
	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

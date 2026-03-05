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
	curso := Curso{
		Nome:         "Go",
		CargaHoraria: 40,
	}

	temp := template.New("CursoTemplate")
	temp, err := temp.Parse("Curso: {{.Nome}}, Carga Horária {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

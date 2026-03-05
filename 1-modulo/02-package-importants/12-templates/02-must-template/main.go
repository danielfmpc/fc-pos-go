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
	temp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}}, Carga Horária {{.CargaHoraria}}"))
	err := temp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

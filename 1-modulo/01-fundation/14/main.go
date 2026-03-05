package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func (c Cliente) Apresentar() {
	c.Nome = "Maria"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", c.Nome, c.Idade)
}

func main() {
	cliente := Cliente{
		Nome:  "João",
		Idade: 30,
	}
	cliente.Apresentar()

	fmt.Printf(cliente.Nome)
}

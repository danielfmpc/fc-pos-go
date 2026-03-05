package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c *Client) Desativar() {
	c.Active = false
}

type Person interface {
	Desativar()
}

type Empresa struct {
	Name    string
	Address Address
}

func (e *Empresa) Desativar() {
	// Lógica para desativar a empresa
}

func DesativarEntidade(p Person) {
	p.Desativar()
}

func main() {
	daniel := Client{
		Name:   "Daniel",
		Age:    30,
		Active: true,
		Address: Address{
			Street: "Rua A",
			Number: 123,
			City:   "São Paulo",
		},
	}

	minhaEmpresa := Empresa{
		Name: "Minha Empresa",
		Address: Address{
			Street: "Avenida B",
			Number: 456,
			City:   "Rio de Janeiro",
		},
	}

	DesativarEntidade(&minhaEmpresa)

	fmt.Printf("Nome: %s, idade: %d, Ativo %t\n", daniel.Name, daniel.Age, daniel.Active)
	fmt.Printf("Endereço: %s, %d, %s\n", daniel.Address.Street, daniel.Address.Number, daniel.Address.City)
}

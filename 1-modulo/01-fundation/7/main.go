package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 42
	d string  = "hello"
	e float64 = 3.14
)

func main() {
	salarios := map[string]float64{
		"João":  3000.0,
		"Maria": 4000.0,
		"Pedro": 3500.0,
	}
	fmt.Println(salarios["Pedro"])
	delete(salarios, "Pedro")
	fmt.Println(salarios)

	sal := make(map[string]int)
	sal["João"] = 3000
	sal["Maria"] = 4000
	sal["Pedro"] = 3500
	fmt.Println(sal)

	for nome, salario := range sal {
		fmt.Printf("%s ganha R$%.2f\n", nome, float64(salario))
	}
}

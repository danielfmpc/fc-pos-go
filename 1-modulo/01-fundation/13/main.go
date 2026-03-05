package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func somaPonteiro(a, b *int) int {
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	resultado := soma(minhaVar1, minhaVar2)
	fmt.Printf("A soma de %d e %d é %d\n", minhaVar1, minhaVar2, resultado)

	resultadoPonteiro := somaPonteiro(&minhaVar1, &minhaVar2)
	fmt.Printf("A soma (com ponteiros) de %d e %d é %d\n", minhaVar1, minhaVar2, resultadoPonteiro)
}

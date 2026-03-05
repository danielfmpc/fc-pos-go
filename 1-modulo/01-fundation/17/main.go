package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Any[T any](a T, b T) bool {
	return true
}

func Compare[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	b := map[string]float64{
		"x": 1.5,
		"y": 2.5,
		"z": 3.5,
	}

	c := map[string]MyNumber{
		"p": 5,
		"q": 10,
		"r": 15,
	}

	fmt.Println(Soma(m))
	fmt.Println(Soma(b))
	fmt.Println(Soma(c))
}

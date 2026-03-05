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
	var meuArray [5]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	meuArray[3] = 4
	meuArray[4] = 5
	// fmt.Printf("O tipo E é %T", e)
	fmt.Println(meuArray[len(meuArray)-1])
}

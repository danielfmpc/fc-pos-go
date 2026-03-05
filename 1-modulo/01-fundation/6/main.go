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
	meuArray := []int{10, 20, 30, 40, 50}

	fmt.Printf("len=%d cap=%d %v\n", len(meuArray), cap(meuArray), meuArray)
	fmt.Printf("len=%d cap=%d %v\n", len(meuArray[:0]), cap(meuArray[:0]), meuArray)
	fmt.Printf("len=%d cap=%d %v\n", len(meuArray[:4]), cap(meuArray[:4]), meuArray[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(meuArray[2:]), cap(meuArray[2:]), meuArray[2:])

	meuArray = append(meuArray, 60)

	fmt.Printf("len=%d cap=%d %v\n", len(meuArray[2:]), cap(meuArray[2:]), meuArray[2:])

	fmt.Println(meuArray[len(meuArray)-1])
}

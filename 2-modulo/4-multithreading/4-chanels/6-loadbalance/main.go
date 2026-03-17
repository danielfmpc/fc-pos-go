package main

import (
	"fmt"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerId, x)
	}
}

func main() {
	data := make(chan int)

	for i := 1; i <= 10; i++ {
		go worker(i, data)
	}

	for i := range 100 {
		data <- i
	}
}

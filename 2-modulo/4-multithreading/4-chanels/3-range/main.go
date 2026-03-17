package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("recebido %d\n", x)
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

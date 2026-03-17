package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	reader(ch, &wg)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("recebido %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

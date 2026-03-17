package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
		waitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func() {
		for i := range 5 {
			fmt.Printf("%d: Task %s is ruinning\n", i, "anonymous")
			time.Sleep(time.Second * 1)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}

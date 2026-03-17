package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task("A")
	go task("B")
}

package main

import "fmt"

func main() {
	c := make(chan string)
	go func() {
		c <- "Hello"
	}()
	msg := <-c
	fmt.Println(msg)
}

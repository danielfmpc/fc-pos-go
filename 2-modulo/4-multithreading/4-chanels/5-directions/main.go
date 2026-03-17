package main

import "fmt"

func recebe(msg string, hello chan<- string) {
	hello <- msg
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)

	go recebe("hello", hello)
	ler(hello)

}

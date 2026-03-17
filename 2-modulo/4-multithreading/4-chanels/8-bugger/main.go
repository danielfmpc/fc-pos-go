package main

func main() {
	ch := make(chan string, 1)

	ch <- "message 1"
	ch <- "message 2"

	println(<-ch)
	println(<-ch)
}

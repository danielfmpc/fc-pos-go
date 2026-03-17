package main

func main() {
	forever := make(chan bool)

	go func() {
		for i := range 10 {
			print(i)
		}
		forever <- true
	}()

	<-forever
}

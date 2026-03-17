package main

import (
	"sync/atomic"
	"time"
)

type Message struct {
	Id   uint64
	Body string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var id uint64 = 0
	go func() {
		for {
			time.Sleep(time.Millisecond)
			atomic.AddUint64(&id, 1)
			msg := Message{
				Id:   id,
				Body: "message Coelhao",
			}
			c1 <- msg
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond)
			atomic.AddUint64(&id, 1)
			msg := Message{
				Id:   id,
				Body: "message Kafka",
			}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			println("received", msg.Id, msg.Body)
		case msg := <-c2:
			println("received", msg.Id, msg.Body)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}

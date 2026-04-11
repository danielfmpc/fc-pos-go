package main

import (
	"fmt"

	"github.com/danielfmpc/pos-go-fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgOut := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgOut, "orders", "amq.direct", "")
	for msg := range msgOut {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}

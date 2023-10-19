package main

import (
	"fmt"
	"os"

	rabbitmq "github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitMQ"
	"github.com/Canhassi12/transaction-microsservice/internal/service"
)

func main() {
	ch, _, err := rabbitmq.Connection()

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	
	msgs, err := ch.Consume(
		"orders",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Print("consume error ", err.Error())
	}

	forever := make(chan bool)
	go func() {
		for order := range msgs {
			service.Create_transaction(order)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

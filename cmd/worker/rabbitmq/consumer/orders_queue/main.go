package main

import (
	"fmt"
	"os"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq"
	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/Canhassi12/transaction-microsservice/internal/service"
)

func main() {
	ch, _, err := rabbitmq.Connection()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	db := db.Connect()
	
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
			service.CreateTransaction(db, order)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

package main

import (
	"fmt"
	"os"

	rabbitmq "github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitMQ"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, ctx, err := rabbitmq.Connection()

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	q, err := ch.QueueDeclare(
		"orders",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	

	err = ch.PublishWithContext(ctx,
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp091.Publishing {
		  DeliveryMode: amqp091.Persistent,
		  ContentType:  "text/plain",
		  Body:         []byte("aaaa"),
		})

	if err != nil {
		fmt.Println(err.Error())
	}

    fmt.Println("Successfully Published Message to Queue")
}

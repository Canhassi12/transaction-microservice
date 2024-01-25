package main

import (
	"fmt"
	"os"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq"
)

func main() {
	var qp = rabbitmq.QueueConnection{}
	if err := qp.Connection(); err != nil {
		panic(err.Error())
	}

	msgs, err := qp.Ch.Consume(
		"dlq",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Print("consume error dlx", err.Error())
		os.Exit(1)
	}

	forever := make(chan bool)
	go func() {
		for order := range msgs {
			fmt.Println("DLX QUEUE: ", string(order.Body), " - ", order.Headers["x-death"])

		}
	}()

	fmt.Println(" [*] - Waiting for dlx queue messages")
	<-forever
}

package main

import (
	"fmt"
	"os"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq"
	transactionQueue "github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq/producer/transaction_queue"
	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/Canhassi12/transaction-microsservice/internal/service"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	var qp = rabbitmq.QueueConnection{}
	if err := qp.Connection(); err != nil {
		panic(err.Error())
	}

	db := db.Connect()

	msgs, err := qp.Ch.Consume(
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
		os.Exit(1)
	}

	forever := make(chan bool)
	go func() {
		for order := range msgs {
			status := service.CreateTransaction(db, order)
			if status == "error" {
				attempts, ok := order.Headers["attempts"].(int32)
				if !ok {
					attempts = 1
				}

				if attempts >= 3 {
					println("Error to process order, sending to DLQ")
				}

				err = qp.Ch.PublishWithContext(qp.Ctx,
					"",       // exchange
					"orders", // routing key
					false,    // mandatory
					false,
					amqp091.Publishing{
						DeliveryMode: amqp091.Persistent,
						ContentType:  "application/json",
						Body:         order.Body,
					})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			order.Ack(true)

			transactionQueue.SendTransactionStatus(status, &qp)
		}
	}()

	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

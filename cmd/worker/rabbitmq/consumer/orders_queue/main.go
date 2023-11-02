package main

import (
	"fmt"
	"os"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq"
	transactionQueue "github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq/producer/transaction_queue"
	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/Canhassi12/transaction-microsservice/internal/service"
)

func main() {
	var qp = rabbitmq.QueueConnection{}
	if err := qp.Connection(); err != nil {
		println(err.Error())
		os.Exit(1)
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
	}

	forever := make(chan bool)
	go func() {
		for order := range msgs {
			status := service.CreateTransaction(db, order)

			transactionQueue.SendTransactionStatus(status, &qp)
		}
	}()

	fmt.Println(" [*] - Waiting for messages")
	<-forever
}


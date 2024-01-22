package transactionQueue

import (
	"fmt"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

func SendTransactionStatus(status string, qp *rabbitmq.QueueConnection) {

	err := qp.Ch.PublishWithContext(qp.Ctx,
		"",             // exchange
		"transactions", // routing key
		false,          // mandatory
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         []byte(status),
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully Published Transaction Status to ORDER microservice")
}

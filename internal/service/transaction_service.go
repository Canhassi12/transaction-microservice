package service

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func Create_transaction(order amqp091.Delivery) {
	fmt.Printf("Recieved Message: %s\n", order.Body)
}

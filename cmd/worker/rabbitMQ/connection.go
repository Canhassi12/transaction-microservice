package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func Connection() (*amqp091.Channel, context.Context, error) {
	conn, err := amqp091.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		return nil, nil, fmt.Errorf("failed Initializing Broker Connection %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf(err.Error())
	}
	defer ch.Close()

	return ch, ctx, nil
}
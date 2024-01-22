package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type QueueConnection struct {
	Ch  *amqp091.Channel
	Ctx context.Context
}

func (qp *QueueConnection) Connection() error {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed Initializing Broker Connection aaaa %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed Initializing Broker Channel 1212 %s", err.Error())
	}

	qp.Ch = ch
	qp.Ctx = ctx

	// defer ch.Close()
	_, err = ch.QueueDeclare(
		"transactions",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("fail declare %s", err.Error())
	}

	_, err = ch.QueueDeclare(
		"orders",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("fail declare 2 %s", err.Error())
	}

	return err
}

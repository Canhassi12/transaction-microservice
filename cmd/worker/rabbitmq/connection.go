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
	conn, err := amqp091.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		return fmt.Errorf("failed Initializing Broker Connection %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	// defer ch.Close()

	qp.Ch = ch
	qp.Ctx = ctx

	return err
}

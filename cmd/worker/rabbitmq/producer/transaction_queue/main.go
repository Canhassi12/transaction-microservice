package transaction_queue

import (
	"github.com/Canhassi12/transaction-microsservice/db"
)

func Send_transaction(t db.Transaction) {
	// conn, err := amqp091.Dial("amqp://user:password@localhost:5672/")
	// if err != nil {
	// 	fmt.Println("failed Initializing Broker Connection", err.Error())
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// q, err := ch.QueueDeclare(
	// 	"transactions",
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// orderJSON, err := json.Marshal(order)
	// if err != nil {
	// 	fmt.Println("Error serializing order:", err)
	// 	return
	// }
	
	// err = ch.PublishWithContext(ctx,
	// 	"",           // exchange
	// 	q.Name,       // routing key
	// 	false,        // mandatory
	// 	false,
	// 	amqp091.Publishing {
	// 	  DeliveryMode: amqp091.Persistent,
	// 	  ContentType:  "application/json",
	// 	  Body:         orderJSON,
	// 	})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

    // fmt.Println("Successfully Published Message to Queue")
}
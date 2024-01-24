package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("failed Initializing Broker Connection", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
	}
	// defer ch.Close()

	address := db.Address{
		Street:        "Rua canhas",
		StreetNumber:  12,
		Neighbourhood: "Logo ali",
		District:      "Aqui que eh aqui?",
		City:          "São Paulo",
		State:         "São Paulo",
		Country:       "Brazil",
		Zipcode:       "95275971",
	}

	order := db.Order{
		ID:             2,
		Status:         "pending",
		UserId:         1,
		Amount:         200,
		PaymentType:    "credit_card",
		DocumentNumber: "80704129094",
		Address:        address,
		Phone:          "5511971178901",
		FullName:       "É O Canhas",
		Email:          "canhassi@gmail.com",
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error serializing order:", err)
		return
	}

	err = ch.PublishWithContext(ctx,
		"",       // exchange
		"orders", // routing key
		false,    // mandatory
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         orderJSON,
		})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully Published Message to Queue")
}

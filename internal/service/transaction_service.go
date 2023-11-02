package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Canhassi12/transaction-microsservice/cmd/worker/rabbitmq/producer/transaction_queue"
	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rabbitmq/amqp091-go"
)

func CreateTransaction(q *sqlx.DB, order amqp091.Delivery) {
	var receivedOrder db.Order
	if err := json.Unmarshal(order.Body, &receivedOrder); err != nil {
		fmt.Println("Error deserializing order:", err)
		return
	}

	jsonData, err := json.Marshal(receivedOrder)
    if err != nil {
        fmt.Printf("Erro ao serializar a struct: %v\n", err)
        return
    }

	payload := bytes.NewBuffer([]byte(jsonData))
	payment := processPayment(payload)

	var t = db.Transaction{}
	result := q.QueryRowx(`INSERT INTO transactions (status, user_id, order_id, paid_at, payment_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, status, user_id, order_id, paid_at, payment_id`, "paid", receivedOrder.UserId, receivedOrder.ID, payment["paid_at"], payment["id"]).StructScan(&t)
	if result != nil { 
		panic("erro insert f")
	}

	transaction_queue.Send_transaction(t)
}

func processPayment(payload *bytes.Buffer) map[string]interface{} {
	// resp, err := http.Post("https://run.mocky.io/v3/8fafdd68-a090-496f-8c9a-3442cf30dae6", "application/json", payload)
	// if err != nil {
	// 	fmt.Printf("Erro ao fazer a requisição: %v\n", err)
	// }
	// defer resp.Body.Close()

	// _, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Erro ao ler a resposta: %v\n", err)
	// }

	id := uuid.NewString()
	
	return map[string]interface{} {
		"id": id,
		"paid_at": time.Now(),
	};
}

package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/rabbitmq/amqp091-go"
)

func Create_transaction(q *sql.DB, order amqp091.Delivery) {
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

	payment := process_payment(payload)

	id, _ := exec.Command("uuidgen").Output() 
	
	_, err = q.Exec("INSERT INTO transactions (id, status, user_id, order_id, paid_at, payment_id) VALUES(?, ?, ?, ?, ?, ?)", id, "paid", receivedOrder.UserId, receivedOrder.ID, payment["paid_at"], payment["id"])
	if err != nil {
		panic(err)
	}

	var t = db.Transaction{}
	row := q.QueryRow("SELECT id, status, user_id, order_id, paid_at, payment_id FROM transactions WHERE id = ?", id)
	if err := row.Scan(&t.ID, &t.Status, &t.UserID, &t.OrderID, &t.PaidAt, &t.PaymentId); err != nil {
		if err != sql.ErrNoRows {
			panic("aaaa") // TODO AMANHAS
		}
	}

	fmt.Printf("ID: %s, Status: %s, UserID: %d, OrderID: %d, PaidAt: %s, PaymentId: %s\n", t.ID, t.Status, t.UserID, t.OrderID, t.PaidAt, t.PaymentId)


	// transaction_queue.Send_transaction(t)
}

func process_payment(payload *bytes.Buffer) map[string]interface{} {
	// resp, err := http.Post("https://run.mocky.io/v3/8fafdd68-a090-496f-8c9a-3442cf30dae6", "application/json", payload)
	// if err != nil {
	// 	fmt.Printf("Erro ao fazer a requisição: %v\n", err)
	// }
	// defer resp.Body.Close()

	// _, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Erro ao ler a resposta: %v\n", err)
	// }

	id, _ := exec.Command("uuidgen").Output() 

	return map[string]interface{} {
		"id": id,
		"paid_at": time.Now(),
	};
}

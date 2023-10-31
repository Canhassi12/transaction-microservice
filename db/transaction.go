package db

import "time"

type Transaction struct {
    ID         string    `json:"id"`
    Status     string    `json:"status"`
    UserID     int       `json:"user_id"`
    OrderID    int       `json:"order_id"`
    PaidAt     time.Time `json:"paid_at"`
    PaymentId  string    `json:"payment_id"`
}

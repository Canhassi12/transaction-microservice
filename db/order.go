package db

type Order struct {
	ID             int     `json:"id"`
	Status         string  `json:"status"`
	UserId         int     `json:"user_id"`
	Amount         float64 `json:"amount"`
	PaymentType    string  `json:"payment_type"`
	DocumentNumber string  `json:"document_number"`
	Address        Address `json:"address"`
	Phone          string  `json:"phone"`
	FullName       string  `json:"full_name"`
	Email          string  `json:"email"`
}

type Address struct {
	Street        string `json:"street"`
	StreetNumber  int    `json:"street_number"`
	Neighbourhood string `json:"neighbourhood"`
	District      string `json:"district"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Zipcode       string `json:"zipcode"`
}
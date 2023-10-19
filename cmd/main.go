package main

import "github.com/Canhassi12/transaction-microsservice/api/handler"


func main() {
	r := handler.RegisterRoutes()
    r.Run(":8080")
}

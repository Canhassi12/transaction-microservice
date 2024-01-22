package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=canhas dbname=tmicroservice password=root host=localhost sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}

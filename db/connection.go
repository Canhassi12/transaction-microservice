package db

import (
	"database/sql"
	"time"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/microsservices") 
		
	if err != nil {
		panic("f na conexão")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Canhassi12/transaction-microsservice/api/handler"
	"github.com/Canhassi12/transaction-microsservice/db"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := handler.RegisterRoutes()

	db := db.Connect()
	if err := executeMigrateFile(db); err != nil {
		println(err.Error())
	}

	r.Run(":8080")
}

func executeMigrateFile(db *sqlx.DB) error {
	content, err := os.ReadFile("./db/migrate.sql")
	if err != nil {
		return fmt.Errorf("error to read migrate file: %w", err)
	}

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		query = strings.Trim(query, " \r\n")
		if query == "" {
			continue
		}

		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query: %s, error: %w", query, err)
		}

	}
	return nil
}

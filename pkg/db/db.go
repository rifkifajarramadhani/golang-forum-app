package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(datasource string) (*sql.DB, error) {
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
		return nil, err
	}
	return db, nil
}
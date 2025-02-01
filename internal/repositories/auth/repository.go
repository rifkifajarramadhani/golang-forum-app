package auth

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Println("query error", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println("scan error", err)
		}
		log.Println("id:", id, "email:", email)
	}

	return &repository{db: db}
}
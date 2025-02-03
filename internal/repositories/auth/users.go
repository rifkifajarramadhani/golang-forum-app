package auth

import (
	"context"
	"database/sql"

	"github.com/rifkifajarramadhani/internal/models/auth"
)

func (r *Repository) CreateUser(ctx context.Context, user auth.UserModel) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(ctx context.Context, email string, password string) (*auth.UserModel, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = ?`
	row := r.db.QueryRowContext(ctx, query, email)

	var user auth.UserModel
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

package auth

import (
	"context"
	"database/sql"

	authModel "github.com/rifkifajarramadhani/internal/models/auth"
)

func (r *repository) CreateUser(ctx context.Context, user authModel.UserModel) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(ctx context.Context, email string, password string) (*authModel.UserModel, error) {
	query := `SELECT email, password FROM users WHERE email = ? AND password = ?`
	row := r.db.QueryRowContext(ctx, query, email, password)

	var user authModel.UserModel
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

package auth

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/auth"
)

type authRepository interface {
	CreateUser(ctx context.Context, model auth.UserModel) error
	GetUser(ctx context.Context, email string, password string) (*auth.UserModel, error)
}

type service struct {
	authRepo authRepository
}

func NewService(authRepo authRepository) *service {
	return &service{
		authRepo: authRepo,
	}
}

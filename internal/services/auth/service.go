package auth

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/auth"
)

type authRepositoryInterface interface {
	CreateUser(ctx context.Context, model auth.UserModel) error
	GetUser(ctx context.Context, email string, password string) (*auth.UserModel, error)
}

type service struct {
	authRepository authRepositoryInterface
}

func NewService(authRepository authRepositoryInterface) *service {
	return &service{
		authRepository: authRepository,
	}
}

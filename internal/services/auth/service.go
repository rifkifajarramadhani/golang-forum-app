package auth

import (
	"context"

	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

type authRepositoryInterface interface {
	CreateUser(ctx context.Context, model auth.UserModel) error
	GetUser(ctx context.Context, email string, password string) (*auth.UserModel, error)
}

type service struct {
	config         *configs.Config
	authRepository authRepositoryInterface
}

func NewService(config *configs.Config, authRepository authRepositoryInterface) *service {
	return &service{
		config:         config,
		authRepository: authRepository,
	}
}

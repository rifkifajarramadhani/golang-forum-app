package auth

import (
	"context"

	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

type AuthRepositoryInterface interface {
	CreateUser(ctx context.Context, model auth.UserModel) error
	GetUser(ctx context.Context, email string, password string) (*auth.UserModel, error)
}

type Service struct {
	config         *configs.Config
	authRepository AuthRepositoryInterface
}

func NewService(config *configs.Config, authRepository AuthRepositoryInterface) *Service {
	return &Service{
		config:         config,
		authRepository: authRepository,
	}
}

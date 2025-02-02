package auth

import (
	"context"
	"errors"

	authModel "github.com/rifkifajarramadhani/internal/models/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, req authModel.RegisterRequest) error {
	user, err := s.authRepository.GetUser(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("email already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	model := authModel.UserModel{
		Username: req.Username,
		Email:    req.Email,
		Password: string(password),
	}

	return s.authRepository.CreateUser(ctx, model)
}

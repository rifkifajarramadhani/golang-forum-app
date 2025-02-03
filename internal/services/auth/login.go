package auth

import (
	"context"
	"errors"

	"github.com/rifkifajarramadhani/internal/models/auth"
	"github.com/rifkifajarramadhani/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req auth.LoginRequest) (string, error) {
	user, err := s.authRepository.GetUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.config.Service.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

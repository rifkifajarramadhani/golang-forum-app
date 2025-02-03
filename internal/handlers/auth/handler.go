package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, req auth.RegisterRequest) error
	Login(ctx context.Context, req auth.LoginRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	authService AuthServiceInterface
}

func NewHandler(api *gin.Engine, authService AuthServiceInterface) *Handler {
	return &Handler{
		Engine:      api,
		authService: authService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/auth")
	route.POST("/register", h.Register)
	route.POST("/login", h.Login)
}

package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

type authService interface {
	Register(ctx context.Context, req auth.RegisterRequest) error
}

type Handler struct {
	*gin.Engine

	authServ authService
}

func NewHandler(api *gin.Engine, authServ authService) *Handler {
	return &Handler{
		Engine:   api,
		authServ: authServ,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/auth")
	route.POST("/register", h.Register)
}

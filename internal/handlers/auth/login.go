package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request auth.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.authService.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := auth.LoginResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, response)
}

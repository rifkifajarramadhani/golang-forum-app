package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/models/auth"
)

func (h *Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var request auth.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.authService.Register(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

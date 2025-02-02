package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secret := configs.GetConfig().Service.JWTSecret
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing authorization header"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secret)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}

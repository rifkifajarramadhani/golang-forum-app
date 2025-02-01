package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/handlers/auth"
)

func main() {
	r := gin.Default()

	authHandler := auth.NewHandler(r)
	authHandler.Register()

	r.Run()
}
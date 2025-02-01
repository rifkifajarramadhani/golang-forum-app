package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/handler/membership"
)

func main() {
	r := gin.Default()

	membershipHandler := membership.NewHandler(r)
	membershipHandler.Register()

	r.Run()
}
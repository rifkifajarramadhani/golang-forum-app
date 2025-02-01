package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.Register()

	r.Run()
}
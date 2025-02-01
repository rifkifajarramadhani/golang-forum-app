package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/internal/handlers/auth"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithCOnfigFolders(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to init configs: %v", err)
	}
	cfg = configs.GetConfig()
	fmt.Printf("config: %+v\n", cfg)

	authHandler := auth.NewHandler(r)
	authHandler.Register()

	r.Run(cfg.Service.Port)
}
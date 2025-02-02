package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/internal/handlers/auth"
	authRepo "github.com/rifkifajarramadhani/internal/repositories/auth"
	authServ "github.com/rifkifajarramadhani/internal/services/auth"
	"github.com/rifkifajarramadhani/pkg/db"
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

	db, err := db.Connect(cfg.Database.Datasource)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	authRepo := authRepo.NewRepository(db)

	authServ := authServ.NewService(authRepo)

	authHandler := auth.NewHandler(r, authServ)
	authHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}

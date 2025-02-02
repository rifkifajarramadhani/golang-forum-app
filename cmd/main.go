package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/configs"
	"github.com/rifkifajarramadhani/internal/handlers/auth"
	"github.com/rifkifajarramadhani/internal/handlers/post"
	authRepository "github.com/rifkifajarramadhani/internal/repositories/auth"
	postRepository "github.com/rifkifajarramadhani/internal/repositories/post"
	authService "github.com/rifkifajarramadhani/internal/services/auth"
	postService "github.com/rifkifajarramadhani/internal/services/post"
	"github.com/rifkifajarramadhani/pkg/db"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders(
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
	authRepo := authRepository.NewRepository(db)
	postRepo := postRepository.NewRepository(db)

	authServ := authService.NewService(cfg, authRepo)
	postServ := postService.NewService(postRepo)

	authHandler := auth.NewHandler(r, authServ)
	authHandler.RegisterRoute()

	postHandler := post.NewHandler(r, postServ)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}

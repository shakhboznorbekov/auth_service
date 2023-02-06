package main

import (
	"context"
	"log"


	"github.com/auth_service/api"
	"github.com/auth_service/config"
	"github.com/auth_service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.CloseDB()

	api.SetUpApi(&cfg,r, storage)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}

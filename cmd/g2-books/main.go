package main

import (
	"github.com/RenzoFudo/http_salem_repo/internal/config"
	"github.com/RenzoFudo/http_salem_repo/internal/server"
	"github.com/RenzoFudo/http_salem_repo/internal/storage"
	"log"
)

func main() {
	cfg := config.ReadConfig()
	log.Println(cfg)
	storage := storage.New()
	server := server.New(cfg.Host, storage)

	if err := server.Run(); err != nil {
		panic(err)
	}
}

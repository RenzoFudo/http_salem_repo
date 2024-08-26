package main

import (
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

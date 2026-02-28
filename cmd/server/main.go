// temporary
package main

import (
	"log"

	"github.com/Reza-1988/go-url-shorten/internal/config"
	"github.com/Reza-1988/go-url-shorten/internal/repository/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = postgres.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connection: ok")
}

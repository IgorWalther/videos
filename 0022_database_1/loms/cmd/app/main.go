package main

import (
	"log"

	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/app"
	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can not initialize logger: %s", err)
	}

	cfg, err := config.New()

	if err != nil {
		log.Fatalf("can not initialize config: %s", err)
	}

	app.Run(logger, cfg)
}

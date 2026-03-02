package main

import (
	"fmt"

	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/app"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		fmt.Println(err.Error()) // todo
		log.Fatalf("can not initialize logger: %s", err)
	}

	app.Run(logger)
}

package controller

import (
	"github.com/igoroutine-courses/the_nature_of_microservices/orders/internal/repository"
	"go.uber.org/zap"
)

type api struct {
	logger           *zap.Logger
	ordersRepository repository.OrdersRepository
}

func New(
	logger *zap.Logger,
	ordersRepository repository.OrdersRepository,
) *api {
	return &api{
		logger:           logger,
		ordersRepository: ordersRepository,
	}
}

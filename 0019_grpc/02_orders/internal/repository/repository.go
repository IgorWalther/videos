package repository

import (
	"context"

	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
)

type (
	OrdersRepository interface {
		CreateOrder(ctx context.Context, order *ordersv1.Order) (*ordersv1.Order, error)
		GetOrder(ctx context.Context, orderID string) (*ordersv1.Order, error)
		ListUserOrders(ctx context.Context, userID string) (<-chan *ordersv1.Order, <-chan error)
	}
)

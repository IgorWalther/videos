package repository

import (
	"context"
	"errors"
	"sync"

	ordersv1 "github.com/igoroutine-courses/the_nature_of_microservices/orders/pkg/api/v1"
)

// TODO: clean architecture

var _ OrdersRepository = (*inMemoryOrdersRepository)(nil)

var (
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrOrderNotFound      = errors.New("order not found")
)

type inMemoryOrdersRepository struct {
	mx     *sync.RWMutex
	orders map[string]*ordersv1.Order
}

func NewInMemoryOrdersRepository() *inMemoryOrdersRepository {
	return &inMemoryOrdersRepository{
		mx:     new(sync.RWMutex),
		orders: make(map[string]*ordersv1.Order),
	}
}

func (r *inMemoryOrdersRepository) CreateOrder(_ context.Context, order *ordersv1.Order) (*ordersv1.Order, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.orders[order.GetId()]; ok {
		return nil, ErrOrderAlreadyExists
	}

	o := order
	r.orders[o.GetId()] = o

	return order, nil
}

func (r *inMemoryOrdersRepository) GetOrder(_ context.Context, orderID string) (*ordersv1.Order, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	v, ok := r.orders[orderID]
	if !ok {
		return nil, ErrOrderNotFound
	}

	return v, nil
}

func (r *inMemoryOrdersRepository) ListUserOrders(
	ctx context.Context,
	userID string,
) (<-chan *ordersv1.Order, <-chan error) {
	out := make(chan *ordersv1.Order)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		r.mx.RLock()
		defer r.mx.RUnlock() // TODO: lock granularity

		for _, o := range r.orders {
			if o.GetUserId() == userID {
				select {
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				case out <- o:
				}
			}
		}
	}()

	return out, errCh
}

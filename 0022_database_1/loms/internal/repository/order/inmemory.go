package order

import (
	"context"
	"slices"
	"sync"
	"time"

	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/entity"
)

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		orders: make(map[int64]*entity.Order),
		nextID: 1,
		stocks: make(map[uint32]uint64),
	}
}

type inMemoryRepository struct {
	mx     sync.RWMutex
	orders map[int64]*entity.Order
	nextID int64
	stocks map[uint32]uint64
}

func (r *inMemoryRepository) nextOrderID() int64 {
	id := r.nextID
	r.nextID++
	return id
}

func (r *inMemoryRepository) CreateOrder(_ context.Context, o *entity.Order) (int64, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	o.ID = r.nextOrderID()
	r.orders[o.ID] = o

	return o.ID, nil
}

func (r *inMemoryRepository) GetOrder(_ context.Context, orderID int64) (*entity.Order, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	o, ok := r.orders[orderID]

	if !ok || o == nil {
		return nil, entity.ErrOrderNotFound
	}

	cp := *o
	cp.Items = slices.Clone(o.Items)

	return &cp, nil
}

func (r *inMemoryRepository) SetOrderStatus(
	_ context.Context,
	orderID int64,
	status entity.OrderStatus,
) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	o, ok := r.orders[orderID]
	if !ok {
		return nil
	}
	o.Status = status
	o.UpdatedAt = time.Now()
	return nil
}

func (r *inMemoryRepository) ReserveStock(_ context.Context, sku uint32, count uint64) bool {
	r.mx.Lock()
	defer r.mx.Unlock()
	if r.stocks[sku] < count {
		return false
	}

	r.stocks[sku] -= count
	return true
}

func (r *inMemoryRepository) ReserveStocks(_ context.Context, items []entity.OrderItem) bool {
	r.mx.Lock()
	defer r.mx.Unlock()

	for _, it := range items {
		if r.stocks[it.SKU] < uint64(it.Count) {
			return false
		}
	}

	for _, it := range items {
		r.stocks[it.SKU] -= uint64(it.Count)
	}
	return true
}

func (r *inMemoryRepository) ReleaseStock(_ context.Context, sku uint32, count uint64) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.stocks[sku] += count
}

func (r *inMemoryRepository) GetStock(_ context.Context, sku uint32) uint64 {
	r.mx.RLock()
	defer r.mx.RUnlock()
	n, _ := r.stocks[sku]
	return n
}

func (r *inMemoryRepository) SetStock(_ context.Context, sku uint32, count uint64) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.stocks[sku] = count
}

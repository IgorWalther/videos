package product

import (
	"context"
	"sync"

	"github.com/igoroutine-courses/microservices.ecommerce.loms/internal/entity"
)

type inMemoryRepository struct {
	mu      sync.RWMutex
	bySKU   map[uint32]entity.Product
	nextSKU uint32
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		bySKU:   make(map[uint32]entity.Product),
		nextSKU: 1,
	}
}

func (r *inMemoryRepository) GetProductBySKU(_ context.Context, sku uint32) (*entity.Product, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.bySKU[sku]
	if !ok {
		return nil, false
	}
	return &p, true
}

func (r *inMemoryRepository) CreateProduct(_ context.Context, name string, price uint32) uint32 {
	r.mu.Lock()
	defer r.mu.Unlock()
	sku := r.nextSKU
	r.nextSKU++
	r.bySKU[sku] = entity.Product{
		Name:  name,
		Price: price,
	}
	return sku
}

func (r *inMemoryRepository) SetProduct(_ context.Context, sku uint32, p entity.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bySKU[sku] = p
}

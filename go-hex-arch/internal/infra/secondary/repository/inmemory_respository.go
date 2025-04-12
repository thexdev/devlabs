package repository

import (
	"sync"
)

var mu sync.Mutex

type InMemoryRepository struct {
	stacks []any
}

func NewInMemoryRespository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (r *InMemoryRepository) All() []any {
	return r.stacks
}

func (r *InMemoryRepository) Save(record any) (bool, error) {
	mu.Lock()
	defer mu.Unlock()

	r.stacks = append(r.stacks, record)

	return true, nil
}

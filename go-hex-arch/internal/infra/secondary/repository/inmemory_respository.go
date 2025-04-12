package repository

import (
	"gohexarch/internal/app/ports/secondary"
	"sync"
)

var mu sync.Mutex

type InMemoryRepository struct {
	stacks []secondary.Record
}

func NewInMemoryRespository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (r *InMemoryRepository) All() []secondary.Record {
	return r.stacks
}

func (r *InMemoryRepository) Save(record secondary.Record) (bool, error) {
	mu.Lock()
	defer mu.Unlock()

	r.stacks = append(r.stacks, record)

	return true, nil
}

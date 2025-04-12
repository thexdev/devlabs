package mock

import (
	"gohexarch/internal/app/ports/secondary"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) All() []secondary.Record {
	return make([]secondary.Record, 0)
}

func (mr *MockRepository) Save(record secondary.Record) (bool, error) {
	args := mr.Called(record)
	return args.Bool(0), args.Error(1)
}

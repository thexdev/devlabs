package mock

import (
	"fmt"
	"gohexarch/internal/app/ports/secondary"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) All() []secondary.Record {
	fmt.Println("All() called")
	args := mr.Called()
	return args.Get(0).([]secondary.Record)
}

func (mr *MockRepository) Save(record secondary.Record) (bool, error) {
	args := mr.Called(record)
	return args.Bool(0), args.Error(1)
}

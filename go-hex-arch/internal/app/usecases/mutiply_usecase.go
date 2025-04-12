package usecases

import (
	"gohexarch/internal/app/ports/secondary"
	"gohexarch/internal/domain/models"

	"github.com/google/uuid"
)

type MultiplyUseCase struct {
	repo secondary.Repository
}

func NewMultiplyUseCase(repo secondary.Repository) *MultiplyUseCase {
	return &MultiplyUseCase{
		repo: repo,
	}
}

func (uc *MultiplyUseCase) Execute(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)

	result, err := calc.Multiply(a, b)

	expression := secondary.Expression{
		A:        a,
		B:        b,
		Operator: "*",
	}

	record := secondary.Record{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	uc.repo.Save(record)

	return result, err
}

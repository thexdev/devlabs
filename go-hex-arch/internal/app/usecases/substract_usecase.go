package usecases

import (
	"gohexarch/internal/app/ports/secondary"
	"gohexarch/internal/domain/models"
)

type SubstractUseCase struct {
	repo secondary.Repository
}

func NewSubtractUseCase(repo secondary.Repository) *SubstractUseCase {
	return &SubstractUseCase{
		repo: repo,
	}
}

func (uc *SubstractUseCase) Execute(a float32, b float32) (float32, error) {
	calc := new(models.Calculator)

	result, err := calc.Subtract(a, b)

	expression := secondary.Expression{
		A: a,
		B: b,
		Operator: "-",
	}

	record := secondary.Record{
		Expression: expression,
		Result: result,
	}

	uc.repo.Save(record)

	return result, err
}

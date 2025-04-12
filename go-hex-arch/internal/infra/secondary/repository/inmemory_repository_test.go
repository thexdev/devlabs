package repository_test

import (
	"gohexarch/internal/app/ports/secondary"
	"gohexarch/internal/infra/secondary/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test utils
func mockRecordHistory(
	a float32, operator string, b float32, result float32,
) secondary.Record {
	expression := secondary.Expression{
		A:        a,
		B:        b,
		Operator: operator,
	}

	record := secondary.Record{
		Expression: expression,
		Result:     result,
	}

	return record
}

func TestInMemoryRepository(t *testing.T) {
	t.Run("initialize repository without error", func(t *testing.T) {
		repo := repository.NewInMemoryRespository()

		assert.IsType(t, repository.InMemoryRepository{}, *repo)
	})
}

func TestInMemoryRepository_All(t *testing.T) {
	t.Run("returns empty slice when theres no records", func(t *testing.T) {
		repo := repository.NewInMemoryRespository()

		assert.Len(t, repo.All(), 0)
	})

	t.Run("returns 1 calculation history", func(t *testing.T) {
		var (
			leftOperand  = float32(1)
			operator     = "+"
			rightOperand = float32(2)
			result       = float32(3)
		)

		record := mockRecordHistory(leftOperand, operator, rightOperand, result)

		repo := repository.NewInMemoryRespository()

		ok, err := repo.Save(record)

		histories := repo.All()

		assert.NoError(t, err)
		assert.True(t, ok)
		assert.Len(t, histories, 1)
	})
}

func TestInMemoryRepository_Save(t *testing.T) {
	t.Run("save calculation record into stakcs", func(t *testing.T) {
		var (
			leftOperand  = float32(1)
			operator     = "+"
			rightOperand = float32(2)
			result       = float32(3)
		)

		record := mockRecordHistory(leftOperand, operator, rightOperand, result)

		repo := repository.NewInMemoryRespository()

		ok, _ := repo.Save(record)

		assert.True(t, ok)
	})
}

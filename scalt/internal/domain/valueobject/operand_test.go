package valueobject_test

import (
	"scalt/internal/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperand(t *testing.T) {
	t.Run("valid operand", func(t *testing.T) {
		operand := valueobject.NewOperand(10)
		assert.Equal(t, 10.0, operand.Value())
	})
}

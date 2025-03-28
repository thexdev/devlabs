package valueobject_test

import (
	"scalt/internal/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperator(t *testing.T) {
	t.Run("valid operator", func(t *testing.T) {
		operator, err := valueobject.NewOperator("+")
		assert.NoError(t, err)
		assert.Equal(t, operator.Symbol(), "+")
	})

	t.Run("invalid operator", func(t *testing.T) {
		operator, err := valueobject.NewOperator("?")
		assert.Error(t,err)
		assert.Equal(t, operator.Symbol(), "")
	})
}

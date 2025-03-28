package entity

import (
	"scalt/internal/domain/valueobject"

	"github.com/google/uuid"
)

type Operation struct {
	ID           string
	LeftOperand  valueobject.Operand
	Operator     valueobject.Operator
	RightOperand valueobject.Operand
}

func NewOperation(
	leftOperand  valueobject.Operand,
	operator     valueobject.Operator,
	rightOperand valueobject.Operand,
) *Operation {
	return &Operation{
		ID: uuid.NewString(),
		LeftOperand: leftOperand,
		Operator: operator,
		RightOperand: rightOperand,
	}
}

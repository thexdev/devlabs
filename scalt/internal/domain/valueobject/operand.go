package valueobject

type Operand struct {
	value float64
}

func NewOperand(value float64) Operand {
	return Operand{value: value}
}

func (o *Operand) Value() float64 {
	return o.value
}

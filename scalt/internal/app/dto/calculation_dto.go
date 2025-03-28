package dto

type CalculationInput struct {
	LeftOperand  float64 `json:"left_operand"`
	Operator     string  `json:"operator"`
	RightOperand float64 `json:"right_operand"`
}

type CalculationOutput struct {
	Result float64 `json:"result"`
}

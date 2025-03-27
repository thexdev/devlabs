package models

type CalculationRequest struct {
	A        float64 `json:"a"`
	B        float64 `json:"b"`
	Operator string  `json:"operator"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

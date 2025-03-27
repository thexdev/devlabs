package models

type CalculationHistory struct {
	ID        string
	Request   CalculationRequest
	Response  CalculationResponse
	Timestamp int64
}

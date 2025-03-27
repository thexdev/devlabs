package models

type CalculationHistory struct {
	ID        string              `json:"id"`
	Request   CalculationRequest  `json:"request"`
	Response  CalculationResponse `json:"response"`
	Timestamp int64               `json:"timestamp"`
}

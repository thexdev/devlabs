package models

import "time"

type Record struct {
	ID         string
	Expression interface{}
	Result     float32
}

type History struct {
	Records []Record
}

func NewHistory() *History {
	return &History{}
}

func (h *History) All() []Record {
	return h.Records
}

func (h *History) Push(expression interface{}, result float32) bool {
	record := Record{
		ID:         time.Now().String(),
		Expression: expression,
		Result:     result,
	}

	h.Records = append(h.Records, record)

	return true
}

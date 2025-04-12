package http

type Request struct {
	A        float32 `json:"a"`
	B        float32 `json:"b"`
	Operator string  `json:"operator"`
}

type Response struct {
	Expression Request `json:"expression"`
	Result     float32 `json:"result"`
	Error      string  `json:"error,omitempty"`
}

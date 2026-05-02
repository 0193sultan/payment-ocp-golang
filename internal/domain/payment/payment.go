package payment

type PaymentRequest struct {
	Amount float64 `json:"amount"`
	Method string  `json:"method"`
}

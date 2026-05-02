package payment

import (
	"github.com/0193sultan/payment-ocp-golang/internal/factory/payment"
)

type PaymentService struct {
	factory *payment.PaymentFactory
}

func NewPaymentService(f *payment.PaymentFactory) *PaymentService {
	return &PaymentService{factory: f}
}

func (s *PaymentService) Pay(method string, amount float64) (string, error) {
	processor, err := s.factory.Get(method)
	if err != nil {
		return "", err
	}
	return processor.Process(amount)
}

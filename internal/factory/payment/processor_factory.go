package payment

import (
	"errors"

	"github.com/0193sultan/payment-ocp-golang/internal/interfaces"
)

type PaymentFactory struct {
	processors map[string]interfaces.PaymentProcessor
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{
		processors: make(map[string]interfaces.PaymentProcessor),
	}
}

func (f *PaymentFactory) Register(p interfaces.PaymentProcessor) {
	f.processors[p.GetName()] = p
}

func (f *PaymentFactory) Get(name string) (interfaces.PaymentProcessor, error) {
	p, ok := f.processors[name]
	if !ok {
		return nil, errors.New("payment method not supported")
	}
	return p, nil
}

package payment

import "fmt"

type BkashProcessor struct{}

func (b *BkashProcessor) Process(amount float64) (string, error) {
	return fmt.Sprintf("Bkash payment processed: %.2f", amount), nil
}

func (b *BkashProcessor) GetName() string {
	return "bkash"
}

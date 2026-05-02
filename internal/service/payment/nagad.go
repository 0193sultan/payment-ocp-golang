package payment

import "fmt"

type NagadProcessor struct{}

func (n *NagadProcessor) Process(amount float64) (string, error) {
	return fmt.Sprintf("Nagad payment processed: %.2f", amount), nil
}

func (n *NagadProcessor) GetName() string {
	return "nagad"
}

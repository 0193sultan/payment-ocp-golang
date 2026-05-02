package interfaces

type PaymentProcessor interface {
	Process(amount float64) (string, error)
	GetName() string
}

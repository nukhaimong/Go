package payment

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	ApiKey string
}
type Nagad struct {
	apiKey string
}

func (b *Bkash) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash\n", amount)
}

func (n *Nagad) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Nagad", amount)
}

type paymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *paymentService {
	return &paymentService{method: method}
}

func (ps paymentService) Checkout() {
	ps.method.pay(100)
}

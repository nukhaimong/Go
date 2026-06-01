package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
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

func (ps paymentService) checkout() {
	ps.method.pay(100)
}

func main() {
	bkash := Bkash{apiKey: "your-api-key"}
	//nagad := Nagad{apiKey: "your-api-key"}

	// bkashService := paymentService{&b}
	// nagadService := paymentService{&n}
	bkashService := NewPaymentService(&bkash)

	bkashService.checkout()
	//nagadService.checkout()
}

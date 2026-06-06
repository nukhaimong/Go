package main

import (
	"learngo/more-on-interface/payment"
)

func main() {
	bkash := payment.Bkash{ApiKey: "your-api-key"}
	//nagad := payment.Nagad{ApiKey: "your-api-key"}

	// bkashService := paymentService{&b}
	// nagadService := paymentService{&n}
	bkashService := payment.NewPaymentService(&bkash)

	bkashService.Checkout()
	//nagadService.Checkout()
}

package main

import "fmt"

// type payment struct{}

// //open closed principle
// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw := razorpay{}
// 	// razorpayPaymentGw.pay(amount)

// 	stripePaymentGw := stripe{}
// 	stripePaymentGw.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	println("Payment of", amount, "made using Razorpay")
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	println("Payment of", amount, "made using Stripe")
// }

//atleast 1 method should be implemented to implement an interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32,account string)
}
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type BrainTree struct {
	gateway paymenter
}

func (b BrainTree) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Braintree")
}

type fakePayment struct {}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Fake Payment")
}

func (f fakePayment) refund(amount float32, account string) {
	fmt.Println("Refund of", amount, "made using Fake Payment to account", account)
}

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(1000.0)

	//interface
	razorpayPaymentGw := BrainTree{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(1000.0)

	fakePaymentGw := fakePayment{}
	newPayment = payment{
		gateway: fakePaymentGw,
	}
	newPayment.makePayment(2000.0)

}

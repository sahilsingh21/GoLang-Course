package main

import "fmt"

// We don't tell which interface we are implemented goLang will do implesentely (auto)
// only condition is that method attributes will the same
// eg. For paymenter is pay(amount float32) and func (r razorpay) pay(amount float32) {}
// pay(amount float32) should be same

// Other Language we specify which interface we are implemented

// atleast one method
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

//Payment gateway
type payment struct {
	gateway paymenter
}

// Open close Principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripepayPaymentGw := stripe{}
	// stripepayPaymentGw.pay(amount)
	// fakepayment := fakepayment{}
	// fakepayment.pay(amount)
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	// logic to make payment
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using fakepayment", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	// logic to make payment
	fmt.Println("Payment is Refunded From PayPal in account: ", account, "amount is: ", amount)
}

func main() {
	// razorpayPaymentGw := razorpay{}
	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}

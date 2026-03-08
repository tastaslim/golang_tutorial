package structures

type IPaymentProcessor interface {
	MakePayment(amount float32)
}
type PhonePePaymentGateway struct{}
type RazorPayPaymentGateway struct{}
type StripePaymentGateway struct{}

type IPaymentGateway struct {
	paymentProcessor IPaymentProcessor
}

func (p IPaymentGateway) MakePayment(amount float32) {
	p.paymentProcessor.MakePayment(amount)
}

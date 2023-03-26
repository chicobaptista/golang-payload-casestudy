package entities

type PaymentMethod interface {
	PostPayment() error
}

type HoldingPaymentMethod struct {
}

func (pm HoldingPaymentMethod) PostPayment() error {
	return nil
}

type MailPaymentMethod struct {
}

func (pm MailPaymentMethod) PostPayment() error {
	return nil
}

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
	Address string
}

func (pm MailPaymentMethod) PostPayment() error {
	return nil
}

type DirectPaymentMethod struct {
	Agency  string
	Account string
}

func (pm DirectPaymentMethod) PostPayment() error {
	return nil
}

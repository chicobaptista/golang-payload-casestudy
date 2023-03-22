package entities

type PaymentSchedule interface {
	IsPayday() bool
}

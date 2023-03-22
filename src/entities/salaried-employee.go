package entities

type SalariedEmployee struct {
	BaseEmployee
	Salary          float64
	PaymentSchedule PaymentSchedule
}

type SalariedPaymentSchedule struct {
}

func (sch SalariedPaymentSchedule) IsPayday() bool {
	return true
}

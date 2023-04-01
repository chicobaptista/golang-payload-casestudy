package entities

type SalariedEmployee struct {
	BaseEmployee
	Salary          float64
	PaymentSchedule PaymentSchedule
}

func (e SalariedEmployee) GetPayment() float64 {
	return e.Salary
}

func NewSalariedEmployee(id int, name string, address string, salary float64) SalariedEmployee {
	return SalariedEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}}, salary, MonthlyPaymentSchedule{}}
}

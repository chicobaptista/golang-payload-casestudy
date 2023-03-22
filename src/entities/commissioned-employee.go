package entities

type CommissionedEmployee struct {
	BaseEmployee
	Salary          float64
	CommissionRate  float64
	PaymentSchedule PaymentSchedule
}

func NewCommissionedEmployee(id int, name string, address string, salary float64, commissionRate float64) CommissionedEmployee {
	return CommissionedEmployee{BaseEmployee{id, name, address}, salary, commissionRate, MonthlyPaymentSchedule{}}
}

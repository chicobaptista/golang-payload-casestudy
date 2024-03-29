package entities

import "time"

type CommissionedEmployee struct {
	BaseEmployee
	Salary          float64
	CommissionRate  float64
	PaymentSchedule PaymentSchedule
	SaleReceipts    []SaleReceipt
}

type SaleReceipt struct {
	Date   time.Time
	Amount float64
}

func NewCommissionedEmployee(id int, name string, address string, salary float64, commissionRate float64) CommissionedEmployee {
	return CommissionedEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}}, salary, commissionRate, MonthlyPaymentSchedule{}, make([]SaleReceipt, 0)}
}

package entities

import "time"

type CommissionedEmployee struct {
	BaseEmployee
	Salary         float64
	CommissionRate float64
	SaleReceipts   []SaleReceipt
}

type SaleReceipt struct {
	Date   time.Time
	Amount float64
}

func (e CommissionedEmployee) GetPayment() float64 {
	return e.Salary
}

func NewCommissionedEmployee(id int, name string, address string, salary float64, commissionRate float64) CommissionedEmployee {
	return CommissionedEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}, MonthlyPaymentSchedule{}}, salary, commissionRate, make([]SaleReceipt, 0)}
}

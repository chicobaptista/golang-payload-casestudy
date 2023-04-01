package entities

import "errors"

type Employee interface {
	GetId() int
	GetPayment() float64
}

type BaseEmployee struct {
	Id            int
	Name          string
	Address       string
	PaymentMethod PaymentMethod
	Affiliation   EmployeeAffiliation
}

func (e BaseEmployee) GetId() int {
	return e.Id
}
func (e BaseEmployee) GetPayment() float64 {
	return 0
}

type EmployeeAffiliation interface {
	GetAffiliationId() (int, error)
}

type NullAffiliation struct{}

func (a NullAffiliation) GetAffiliationId() (int, error) {
	return 0, errors.New("Employee is not Affiliated to any Organization.")
}

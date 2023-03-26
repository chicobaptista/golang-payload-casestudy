package entities

import "errors"

type Employee interface {
	GetId() int
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

type EmployeeAffiliation interface {
	GetAffiliationId() (int, error)
}

type NullAffiliation struct{}

func (a NullAffiliation) GetAffiliationId() (int, error) {
	return 0, errors.New("Employee is not Affiliated to any Organization.")
}

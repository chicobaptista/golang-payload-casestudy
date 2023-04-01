package entities

import (
	"errors"
	"time"
)

type Employee interface {
	GetId() int
	GetPayment() float64
	IsPayday(date time.Time) bool
}

type BaseEmployee struct {
	Id              int
	Name            string
	Address         string
	PaymentMethod   PaymentMethod
	Affiliation     EmployeeAffiliation
	PaymentSchedule PaymentSchedule
}

func (e BaseEmployee) GetId() int {
	return e.Id
}
func (e BaseEmployee) GetPayment() float64 {
	return 0
}
func (e BaseEmployee) IsPayday(date time.Time) bool {
	return e.PaymentSchedule.IsPayday(date)
}

type EmployeeAffiliation interface {
	GetAffiliationId() (int, error)
}

type NullAffiliation struct{}

func (a NullAffiliation) GetAffiliationId() (int, error) {
	return 0, errors.New("Employee is not Affiliated to any Organization.")
}

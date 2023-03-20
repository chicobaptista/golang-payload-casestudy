package usecases

import (
	"chicobaptista.github.com/entities"
)

type AddSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float64
	eRepo   EmployeeRepository
}

func (tx AddSalariedEmployee) Execute() (success bool, err error) {
	e := entities.SalariedEmployee{tx.Salary, entities.BaseEmployee{tx.Id, tx.Name, tx.Address}}
	tx.eRepo.AddEmployee(e)
	return true, nil
}

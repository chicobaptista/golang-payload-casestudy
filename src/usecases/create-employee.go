package usecases

import (
	"chicobaptista.github.com/entities"
)

type AddSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float32
	eRepo   EmployeeRepository
}

func (tx AddSalariedEmployee) Execute() (success bool, err error) {
	tx.eRepo.AddEmployee(entities.Employee{tx.Id, tx.Name})
	return true, nil
}

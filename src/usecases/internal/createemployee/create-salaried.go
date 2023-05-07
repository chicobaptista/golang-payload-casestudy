package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type CreateSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float64
	eRepo   interfaces.EmployeeRepository
}

func (tx CreateSalariedEmployee) Execute() (success bool, err error) {
	e := entities.NewSalariedEmployee(tx.Id, tx.Name, tx.Address, tx.Salary)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

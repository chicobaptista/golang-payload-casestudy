package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeToSalaried struct {
	Id     int
	Salary float64
	eRepo  EmployeeRepository
}

func (tx ChangeEmployeeToSalaried) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	se := entities.NewSalariedEmployee(be.Id, be.Name, be.Address, tx.Salary)

	tx.eRepo.AddEmployee(se)
	return true, nil
}

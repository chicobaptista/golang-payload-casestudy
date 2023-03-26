package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeName struct {
	Id    int
	Name  string
	eRepo EmployeeRepository
}

func (tx ChangeEmployeeName) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, ok := e.(entities.BaseEmployee)

	be.Name = tx.Name

	tx.eRepo.AddEmployee(be)
	return true, nil
}

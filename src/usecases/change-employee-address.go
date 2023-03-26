package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeAddress struct {
	Id      int
	Address string
	eRepo   EmployeeRepository
}

func (tx ChangeEmployeeAddress) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, ok := e.(entities.BaseEmployee)

	be.Address = tx.Address

	tx.eRepo.AddEmployee(be)
	return true, nil
}

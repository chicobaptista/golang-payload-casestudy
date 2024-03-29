package usecases

import (
	"errors"
	"fmt"
)

type DeleteEmployee struct {
	Id    int
	eRepo EmployeeRepository
}

func (tx DeleteEmployee) Execute() (bool, error) {
	_, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	tx.eRepo.DeleteEmployee(tx.Id)
	return true, nil
}

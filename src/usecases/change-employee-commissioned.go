package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeToCommissioned struct {
	Id             int
	Salary         float64
	CommissionRate float64
	eRepo          EmployeeRepository
}

func (tx ChangeEmployeeToCommissioned) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	ce := entities.NewCommissionedEmployee(be.Id, be.Name, be.Address, tx.Salary, tx.CommissionRate)

	tx.eRepo.AddEmployee(ce)
	return true, nil
}

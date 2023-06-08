package changeemployee

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/payrollcasestudy/entities"
	"chicobaptista.github.com/payrollcasestudy/usecases/interfaces"
)

type ChangeEmployeeToHolding struct {
	Id    int
	eRepo interfaces.EmployeeRepository
}

func (tx ChangeEmployeeToHolding) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	be.PaymentMethod = entities.HoldingPaymentMethod{}

	tx.eRepo.AddEmployee(be)
	return true, nil
}

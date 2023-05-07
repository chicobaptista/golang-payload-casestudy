package changeemployee

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type ChangeEmployeeToDirect struct {
	Id      int
	Agency  string
	Account string
	eRepo   interfaces.EmployeeRepository
}

func (tx ChangeEmployeeToDirect) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	be.PaymentMethod = entities.DirectPaymentMethod{Agency: tx.Agency, Account: tx.Account}

	tx.eRepo.AddEmployee(be)
	return true, nil
}

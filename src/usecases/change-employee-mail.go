package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeToMail struct {
	Id      int
	Address string
	eRepo   EmployeeRepository
}

func (tx ChangeEmployeeToMail) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	be.PaymentMethod = entities.MailPaymentMethod{tx.Address}

	tx.eRepo.AddEmployee(be)
	return true, nil
}

package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type ChangeEmployeeToMail struct {
	Id      int
	Address string
	eRepo   interfaces.EmployeeRepository
}

func (tx ChangeEmployeeToMail) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	be.PaymentMethod = entities.MailPaymentMethod{Address: tx.Address}

	tx.eRepo.AddEmployee(be)
	return true, nil
}

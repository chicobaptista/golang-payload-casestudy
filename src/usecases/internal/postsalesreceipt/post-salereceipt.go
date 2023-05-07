package postsalesreceipt

import (
	"errors"
	"fmt"
	"time"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type PostSaleReceipt struct {
	Id     int
	Date   time.Time
	Amount float64
	eRepo  interfaces.EmployeeRepository
}

func (tx PostSaleReceipt) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))

	}
	ce, ok := e.(entities.CommissionedEmployee)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d is not a Commissioned Employee`, tx.Id))
	}
	ce.SaleReceipts = append(ce.SaleReceipts, entities.SaleReceipt{Date: tx.Date, Amount: tx.Amount})
	tx.eRepo.AddEmployee(ce)
	return true, nil
}

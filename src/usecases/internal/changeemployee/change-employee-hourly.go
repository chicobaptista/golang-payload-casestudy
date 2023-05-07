package changeemployee

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type ChangeEmployeeToHourly struct {
	Id         int
	HourlyRate float64
	eRepo      interfaces.EmployeeRepository
}

func (tx ChangeEmployeeToHourly) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	he := entities.NewHourlyEmployee(be.Id, be.Name, be.Address, tx.HourlyRate)

	tx.eRepo.AddEmployee(he)
	return true, nil
}

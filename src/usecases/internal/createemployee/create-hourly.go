package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type CreateHourlyEmployee struct {
	Id         int
	Name       string
	Address    string
	HourlyRate float64
	eRepo      interfaces.EmployeeRepository
}

func (tx CreateHourlyEmployee) Execute() (success bool, err error) {
	e := entities.NewHourlyEmployee(tx.Id, tx.Name, tx.Address, tx.HourlyRate)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

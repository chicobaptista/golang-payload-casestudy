package createemployee

import (
	"chicobaptista.github.com/payrollcasestudy/entities"
)

type CreateHourlyEmployee struct {
	Id         int
	Name       string
	Address    string
	HourlyRate float64
}

func (ceBhv CreateHourlyEmployee) generateEmployee() entities.Employee {
	return entities.NewHourlyEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.HourlyRate)
}

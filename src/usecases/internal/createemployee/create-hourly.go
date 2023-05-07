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

func (ceBhv CreateHourlyEmployee) generateEmployee() entities.Employee {
	return entities.NewHourlyEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.HourlyRate)
}

func (ceBhv CreateHourlyEmployee) saveEmployee(e entities.Employee) {
	ceBhv.eRepo.AddEmployee(e)
}

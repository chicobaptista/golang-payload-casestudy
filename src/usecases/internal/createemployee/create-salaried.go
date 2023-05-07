package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type CreateSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float64
	eRepo   interfaces.EmployeeRepository
}

func (ceBhv CreateSalariedEmployee) generateEmployee() entities.Employee {
	return entities.NewSalariedEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.Salary)
}

func (ceBhv CreateSalariedEmployee) saveEmployee(e entities.Employee) {
	ceBhv.eRepo.AddEmployee(e)
}

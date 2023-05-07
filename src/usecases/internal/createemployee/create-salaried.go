package createemployee

import (
	"chicobaptista.github.com/entities"
)

type CreateSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float64
}

func (ceBhv CreateSalariedEmployee) generateEmployee() entities.Employee {
	return entities.NewSalariedEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.Salary)
}

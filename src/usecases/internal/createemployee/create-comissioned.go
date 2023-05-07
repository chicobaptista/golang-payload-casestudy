package createemployee

import (
	"chicobaptista.github.com/entities"
)

type CreateCommissionedEmployee struct {
	Id             int
	Name           string
	Address        string
	Salary         float64
	CommissionRate float64
}

func (ceBhv CreateCommissionedEmployee) generateEmployee() entities.Employee {
	return entities.NewCommissionedEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.Salary, ceBhv.CommissionRate)
}

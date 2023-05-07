package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type CreateCommissionedEmployee struct {
	Id             int
	Name           string
	Address        string
	Salary         float64
	CommissionRate float64
	eRepo          interfaces.EmployeeRepository
}

func (ceBhv CreateCommissionedEmployee) generateEmployee() entities.Employee {
	return entities.NewCommissionedEmployee(ceBhv.Id, ceBhv.Name, ceBhv.Address, ceBhv.Salary, ceBhv.CommissionRate)
}

func (ceBhv CreateCommissionedEmployee) saveEmployee(e entities.Employee) {
	ceBhv.eRepo.AddEmployee(e)
}

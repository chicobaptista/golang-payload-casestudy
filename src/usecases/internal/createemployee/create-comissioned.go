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

func (tx CreateCommissionedEmployee) Execute() (success bool, err error) {
	e := entities.NewCommissionedEmployee(tx.Id, tx.Name, tx.Address, tx.Salary, tx.CommissionRate)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

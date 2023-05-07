package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type AddSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float64
	eRepo   interfaces.EmployeeRepository
}

func (tx AddSalariedEmployee) Execute() (success bool, err error) {
	e := entities.NewSalariedEmployee(tx.Id, tx.Name, tx.Address, tx.Salary)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

type AddCommissionedEmployee struct {
	Id             int
	Name           string
	Address        string
	Salary         float64
	CommissionRate float64
	eRepo          interfaces.EmployeeRepository
}

func (tx AddCommissionedEmployee) Execute() (success bool, err error) {
	e := entities.NewCommissionedEmployee(tx.Id, tx.Name, tx.Address, tx.Salary, tx.CommissionRate)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

type AddHourlyEmployee struct {
	Id         int
	Name       string
	Address    string
	HourlyRate float64
	eRepo      interfaces.EmployeeRepository
}

func (tx AddHourlyEmployee) Execute() (success bool, err error) {
	e := entities.NewHourlyEmployee(tx.Id, tx.Name, tx.Address, tx.HourlyRate)
	tx.eRepo.AddEmployee(e)
	return true, nil
}

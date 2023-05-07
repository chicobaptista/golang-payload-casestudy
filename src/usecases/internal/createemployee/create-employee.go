package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type CreateEmployeeBehavior interface {
	generateEmployee() entities.Employee
}

type CreateEmployee struct {
	ce    CreateEmployeeBehavior
	eRepo interfaces.EmployeeRepository
}

func (tx CreateEmployee) Execute() (success bool, err error) {
	e := tx.ce.generateEmployee()
	tx.saveEmployee(e)
	return true, nil
}

func (tx CreateEmployee) saveEmployee(e entities.Employee) {
	tx.eRepo.AddEmployee(e)
}

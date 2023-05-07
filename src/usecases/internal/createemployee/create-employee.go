package createemployee

import "chicobaptista.github.com/entities"

type CreateEmployeeBehavior interface {
	generateEmployee() entities.Employee
	saveEmployee(entities.Employee)
}

type CreateEmployee struct {
	ce CreateEmployeeBehavior
}

func (tx CreateEmployee) Execute() (success bool, err error) {
	e := tx.ce.generateEmployee()
	tx.ce.saveEmployee(e)
	return true, nil
}

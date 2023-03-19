package usecases

import "chicobaptista.github.com/entities"

type EmployeeRepository interface {
	GetEmployee(empId int) entities.Employee
	AddEmployee(e entities.Employee)
}

type EmployeeRepositoryImp struct {
	employees map[int]entities.Employee
}

func (er EmployeeRepositoryImp) GetEmployee(empId int) entities.Employee {
	return er.employees[empId]
}

func (er EmployeeRepositoryImp) AddEmployee(e entities.Employee) {
	er.employees[e.Id] = e
}

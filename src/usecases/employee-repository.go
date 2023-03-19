package usecases

import "chicobaptista.github.com/entities"

type EmployeeRepository interface {
	GetEmployee(empId int) entities.Employee
	AddEmployee(e entities.Employee)
}

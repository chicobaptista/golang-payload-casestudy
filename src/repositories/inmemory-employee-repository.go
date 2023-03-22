package repositories

import "chicobaptista.github.com/entities"

type InMemoryEmployeeRepository struct {
	employees map[int]entities.Employee
}

func (er InMemoryEmployeeRepository) GetEmployee(empId int) entities.Employee {
	return er.employees[empId]
}

func (er InMemoryEmployeeRepository) AddEmployee(e entities.Employee) {
	er.employees[e.GetId()] = e
}

func (er InMemoryEmployeeRepository) DeleteEmployee(empId int) {
	delete(er.employees, empId)
}

func MakeInMemoryEmployeeRepository() InMemoryEmployeeRepository {
	return InMemoryEmployeeRepository{make(map[int]entities.Employee)}
}

package usecases

type EmployeeRepository interface {
	GetEmployee(empId int) Employee
	AddEmployee(e Employee)
}

type EmployeeRepositoryImp struct {
	employees map[int]Employee
}

func (er EmployeeRepositoryImp) GetEmployee(empId int) Employee {
	return er.employees[empId]
}

func (er EmployeeRepositoryImp) AddEmployee(e Employee) {
	er.employees[e.Id] = e
}

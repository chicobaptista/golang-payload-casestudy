package usecases

type AddSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float32
	eRepo   EmployeeRepository
}

func (tx AddSalariedEmployee) Execute() (success bool, err error) {
	tx.eRepo.AddEmployee(Employee{tx.Id, tx.Name})
	return true, nil
}

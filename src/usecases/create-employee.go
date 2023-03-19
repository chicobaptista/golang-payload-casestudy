package usecases

type AddSalariedEmployee struct {
	Id      int
	Name    string
	Address string
	Salary  float32
}

func (tx AddSalariedEmployee) Execute() (success bool, err error) {
	return true, nil
}

package entities

type SalariedEmployee struct {
	Salary float64
	BaseEmployee
}

func (e SalariedEmployee) GetId() int {
	return e.Id
}

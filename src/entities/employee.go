package entities

type Employee interface {
	GetId() int
	GetPaymentAmount() float64
}

type BaseEmployee struct {
	Id      int
	Name    string
	Address string
}

func (e BaseEmployee) GetId() int {
	return e.Id
}

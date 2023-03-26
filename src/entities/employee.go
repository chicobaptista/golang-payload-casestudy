package entities

type Employee interface {
	GetId() int
}

type BaseEmployee struct {
	Id            int
	Name          string
	Address       string
	PaymentMethod PaymentMethod
}

func (e BaseEmployee) GetId() int {
	return e.Id
}

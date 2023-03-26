package entities

type UnionMember struct {
	Id      int
	Dues    float64
	Charges []UnionCharge
}

type UnionCharge struct {
	Amount float64
}

func NewUnionMember(id int, dues float64) UnionMember {
	return UnionMember{id, dues, make([]UnionCharge, 0)}
}

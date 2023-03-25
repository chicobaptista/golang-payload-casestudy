package entities

type UnionMember struct {
	Id      int
	Dues    float32
	Charges []UnionCharge
}

type UnionCharge struct {
	Amount float64
}

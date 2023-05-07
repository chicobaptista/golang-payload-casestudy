package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type PostUnionCharge struct {
	MemberId int
	Amount   float64
	empRepo  interfaces.EmployeeRepository
}

func (tx PostUnionCharge) Execute() (bool, error) {
	um, ok := tx.empRepo.GetUnionMember(tx.MemberId)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Member %d not found.`, tx.MemberId))
	}
	um.Charges = append(um.Charges, entities.UnionCharge{Amount: tx.Amount})
	tx.empRepo.PutUnionMember(um)
	return true, nil
}

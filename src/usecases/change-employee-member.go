package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

type ChangeEmployeeToMember struct {
	EmpId    int
	MemberId int
	Dues     float64
	eRepo    interfaces.EmployeeRepository
}

func (tx ChangeEmployeeToMember) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.EmpId)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.EmpId))
	}
	be, _ := e.(entities.BaseEmployee)
	_, ok = be.Affiliation.(entities.NullAffiliation)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d is already a Member`, tx.EmpId))
	}
	_, ok = tx.eRepo.GetUnionMember(tx.MemberId)
	if ok {
		return false, errors.New(fmt.Sprintf(`Member %d is already registered`, tx.MemberId))
	}

	be.Affiliation = entities.UnionAffiliation{Id: tx.MemberId}
	um := entities.NewUnionMember(tx.MemberId, tx.Dues)

	tx.eRepo.PutUnionMember(um)
	tx.eRepo.AddEmployee(be)
	return true, nil
}

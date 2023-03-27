package usecases

import (
	"errors"
	"fmt"

	"chicobaptista.github.com/entities"
)

type ChangeEmployeeRemoveMember struct {
	Id    int
	eRepo EmployeeRepository
}

func (tx ChangeEmployeeRemoveMember) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.Id)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.Id))
	}
	be, _ := e.(entities.BaseEmployee)
	a, ok := be.Affiliation.(entities.NullAffiliation)
	if ok {
		return false, errors.New(fmt.Sprintf(`Employee %d is not a Member`, tx.Id))
	}

	be.Affiliation = entities.NullAffiliation{}
	unionId, _ := a.GetAffiliationId()
	tx.eRepo.DeleteUnionMember(unionId)

	tx.eRepo.AddEmployee(be)
	return true, nil
}

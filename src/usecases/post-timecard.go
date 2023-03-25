package usecases

import (
	"errors"
	"fmt"
	"time"

	"chicobaptista.github.com/entities"
)

type PostTimecard struct {
	EmpId int
	Date  time.Time
	Hours float64
	eRepo EmployeeRepository
}

func (tx PostTimecard) Execute() (bool, error) {
	e, ok := tx.eRepo.GetEmployee(tx.EmpId)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d not found`, tx.EmpId))

	}
	he, ok := e.(entities.HourlyEmployee)
	if !ok {
		return false, errors.New(fmt.Sprintf(`Employee %d is not an Hourly Employee`, tx.EmpId))
	}
	he.Timecards = append(he.Timecards, entities.TimeCard{tx.Date, tx.Hours})
	tx.eRepo.AddEmployee(he)
	return true, nil
}

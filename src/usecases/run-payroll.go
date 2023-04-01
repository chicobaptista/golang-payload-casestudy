package usecases

import "time"

type RunPayroll struct {
	Date    time.Time
	Payroll map[int]Paycheck
	eRepo   EmployeeRepository
}

type Paycheck struct {
	Amount float64
}

func (tx RunPayroll) Execute() (bool, error) {
	empIds := tx.eRepo.GetAllEmployeeIds()
	for _, id := range empIds {
		tx.Payroll[id] = Paycheck{Amount: 0}
	}
	return true, nil
}

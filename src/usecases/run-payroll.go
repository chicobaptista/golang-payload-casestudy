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
		e, _ := tx.eRepo.GetEmployee((id))
		if e.IsPayday(tx.Date) {
			tx.Payroll[id] = Paycheck{Amount: e.GetPayment()}
		}
	}
	return true, nil
}

func (tx RunPayroll) GetPaycheck(empId int) (Paycheck, bool) {
	pc, ok := tx.Payroll[empId]
	return pc, ok
}

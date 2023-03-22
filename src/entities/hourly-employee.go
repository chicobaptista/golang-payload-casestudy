package entities

import "time"

type HourlyEmployee struct {
	BaseEmployee
	HourlyRate      float64
	PaymentSchedule PaymentSchedule
}

func NewHourlyEmployee(id int, name string, address string, hourlyRate float64) HourlyEmployee {
	return HourlyEmployee{BaseEmployee{id, name, address}, hourlyRate, WeeklyPaymentSchedule{}}
}

type WeeklyPaymentSchedule struct {
}

func (sch WeeklyPaymentSchedule) IsPayday(date time.Time) bool {
	return true
}

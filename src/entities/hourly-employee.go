package entities

import (
	"time"
)

type HourlyEmployee struct {
	BaseEmployee
	HourlyRate      float64
	PaymentSchedule PaymentSchedule
	Timecards       []TimeCard
}

type TimeCard struct {
	Date  time.Time
	Hours float64
}

func NewHourlyEmployee(id int, name string, address string, hourlyRate float64) HourlyEmployee {
	return HourlyEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}}, hourlyRate, WeeklyPaymentSchedule{}, make([]TimeCard, 0)}
}

type WeeklyPaymentSchedule struct {
}

func (sch WeeklyPaymentSchedule) IsPayday(date time.Time) bool {
	return true
}

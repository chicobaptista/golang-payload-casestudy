package entities

import (
	"time"
)

type HourlyEmployee struct {
	BaseEmployee
	HourlyRate float64
	Timecards  []TimeCard
}

type TimeCard struct {
	Date  time.Time
	Hours float64
}

func (e HourlyEmployee) GetPayment() float64 {
	hoursWorked := getTotalHoursWorked(e.Timecards)
	return hoursWorked * e.HourlyRate
}

func getTotalHoursWorked(tcs []TimeCard) float64 {
	var hoursWorked float64
	for _, tc := range tcs {
		hoursWorked += tc.Hours
		if tc.Hours > 8.00 {
			hoursWorked += (tc.Hours - 8.00) * 0.5
		}
	}
	return hoursWorked
}

func NewHourlyEmployee(id int, name string, address string, hourlyRate float64) HourlyEmployee {
	return HourlyEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}, WeeklyPaymentSchedule{}}, hourlyRate, make([]TimeCard, 0)}
}

type WeeklyPaymentSchedule struct {
}

func (sch WeeklyPaymentSchedule) IsPayday(date time.Time) bool {
	return true
}

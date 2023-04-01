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

var MAX_WORK_HOURS = 8.00
var OVERTIME_EXTRA_RATE = 0.5

func (e HourlyEmployee) GetPayment() float64 {
	hoursWorked := getTotalHoursWorked(e.Timecards)
	return hoursWorked * e.HourlyRate
}

func getTotalHoursWorked(tcs []TimeCard) float64 {
	var hoursWorked float64
	for _, tc := range tcs {
		hoursWorked += tc.Hours
		hoursWorked += addOvertimeRateIfNeeded(tc.Hours)
	}
	return hoursWorked
}

func addOvertimeRateIfNeeded(hours float64) float64 {
	if hours > MAX_WORK_HOURS {
		return (hours - MAX_WORK_HOURS) * OVERTIME_EXTRA_RATE
	}
	return 0
}

func NewHourlyEmployee(id int, name string, address string, hourlyRate float64) HourlyEmployee {
	return HourlyEmployee{BaseEmployee{id, name, address, HoldingPaymentMethod{}, NullAffiliation{}, WeeklyPaymentSchedule{}}, hourlyRate, make([]TimeCard, 0)}
}

type WeeklyPaymentSchedule struct {
}

func (sch WeeklyPaymentSchedule) IsPayday(date time.Time) bool {
	return true
}

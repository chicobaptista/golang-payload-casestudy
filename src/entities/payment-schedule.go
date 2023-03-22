package entities

import "time"

type PaymentSchedule interface {
	IsPayday(date time.Time) bool
}

type MonthlyPaymentSchedule struct {
}

func (sch MonthlyPaymentSchedule) IsPayday(date time.Time) bool {
	isFriday := date.Weekday() == time.Friday
	SEVEN_DAYS := time.Hour * 24 * 7
	isLastFriday := isFriday && date.Add(SEVEN_DAYS).Month() != date.Month()
	return isLastFriday
}

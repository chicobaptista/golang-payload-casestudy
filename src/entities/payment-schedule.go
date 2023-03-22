package entities

import "time"

type PaymentSchedule interface {
	IsPayday(date time.Time) bool
}

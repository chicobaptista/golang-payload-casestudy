package usecases

import (
	"chicobaptista.github.com/entities"
	"testing"
	"time"
)

func TestIsPaydayForSalariedEmployee(t *testing.T) {
	se := entities.NewSalariedEmployee(1, "Bob", "Home", 1000.00)

	sch, ok := se.PaymentSchedule.(entities.SalariedPaymentSchedule)

	if !ok {
		t.Fatalf(`Expected Schedule to be Salaried`)
	}
	lastFriday := time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local)
	if !sch.IsPayday(lastFriday) {
		t.Fatalf("Salaried Employee should be paid on the last Friday of the month")
	}

	firstFriday := time.Date(2023, 3, 3, 12, 30, 30, 100, time.Local)
	if sch.IsPayday(firstFriday) {
		t.Fatalf("Salaried Employee should be paid on the last Friday of the month")
	}

	lastThursday := time.Date(2023, 3, 30, 12, 30, 30, 100, time.Local)
	if sch.IsPayday(lastThursday) {
		t.Fatalf("Salaried Employee should be paid on the last Friday of the month")
	}
}

package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"math"
	"testing"
	"time"
)

func TestAddSalariedEmployee(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = AddSalariedEmployee{empId, "Bob", "Home", 1000.00, er}

	tx.Execute()

	e := er.GetEmployee(empId)

	se, ok := e.(entities.SalariedEmployee)
	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee instance to be of type Salaried")
	}

	if se.Name != "Bob" {
		t.Fatalf(`Failed to persist Employee Data properly, want Name to be %q, got %v`, "Bob", se.Name)
	}

	if se.Address != "Home" {
		t.Fatalf(`Failed to persist Employee Data properly, want Address to be %q, got %v`, "Home", se.Address)
	}

	if diff := math.Abs(1000.00 - se.Salary); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Salary to be %f, got %f`, 1000.00, se.Salary)
	}

	if diff := math.Abs(1000.00 - se.GetPaymentAmount()); diff > 0.001 {
		t.Fatalf(`Want PaymentAmount to be %f, got %f`, 1000.00, se.GetPaymentAmount())
	}

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

func TestAddCommissionedEmployee(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = AddCommissionedEmployee{empId, "Bob", "Home", 1000.00, 10, er}

	tx.Execute()

	e := er.GetEmployee(empId)

	ce, ok := e.(entities.CommissionedEmployee)
	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee instance to be of type Commissioned")
	}

	if ce.Name != "Bob" {
		t.Fatalf(`Failed to persist Employee Data properly, want Name to be %q, got %v`, "Bob", ce.Name)
	}

	if ce.Address != "Home" {
		t.Fatalf(`Failed to persist Employee Data properly, want Address to be %q, got %v`, "Home", ce.Address)
	}

	if diff := math.Abs(1000.00 - ce.Salary); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Salary to be %f, got %f`, 1000.00, ce.Salary)
	}

	if diff := math.Abs(10.00 - ce.CommissionRate); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Commission Rate to be %f, got %f`, 10.00, ce.CommissionRate)
	}
}

package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"math"
	"testing"
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

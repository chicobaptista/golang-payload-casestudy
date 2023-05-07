package createemployee

import (
	"math"
	"testing"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"chicobaptista.github.com/usecases/interfaces"
)

func TestAddSalariedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx interfaces.Transaction
	tx = CreateEmployee{CreateSalariedEmployee{empId, "Bob", "Home", 1000.00, er}}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

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
}

func TestAddCommissionedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx interfaces.Transaction
	tx = CreateEmployee{CreateCommissionedEmployee{empId, "Bob", "Home", 1000.00, 10, er}}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

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

	if len(ce.SaleReceipts) != 0 {
		t.Fatalf("Should not start with any Sale Receipts associated")
	}
}

func TestAddHourlyEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx interfaces.Transaction
	tx = CreateEmployee{CreateHourlyEmployee{empId, "Bob", "Home", 15.00, er}}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

	he, ok := e.(entities.HourlyEmployee)
	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee instance to be of type Hourly")
	}

	if he.Name != "Bob" {
		t.Fatalf(`Failed to persist Employee Data properly, want Name to be %q, got %v`, "Bob", he.Name)
	}

	if he.Address != "Home" {
		t.Fatalf(`Failed to persist Employee Data properly, want Address to be %q, got %v`, "Home", he.Address)
	}

	if diff := math.Abs(15.00 - he.HourlyRate); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Hourly Rate to be %f, got %f`, 15.00, he.HourlyRate)
	}
}

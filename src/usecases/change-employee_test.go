package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestChangeEmployeeName(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{empId, "Bob", "Home"})

	var tx Transaction
	tx = ChangeEmployeeName{empId, "Jeff", er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)
	be := e.(entities.BaseEmployee)

	if be.Name != "Jeff" {
		t.Fatalf(`Failed to change Employee Name, want %s, got %s`, "Jeff", be.Name)
	}

}

func TestChangeEmployeeNameOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeName{empId, "Jeff", er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the name of a Non-Existing Employee.")
	}

}

func TestChangeEmployeeAddress(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{empId, "Bob", "Home"})

	var tx Transaction
	tx = ChangeEmployeeAddress{empId, "Work", er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)
	be := e.(entities.BaseEmployee)

	if be.Address != "Work" {
		t.Fatalf(`Failed to change Employee Address, want %s, got %s`, "Work", be.Address)
	}

}

func TestChangeEmployeeAddressOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeAddress{empId, "Work", er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the address of a Non-Existing Employee.")
	}

}

func TestChangeEmployeeToHourly(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{empId, "Bob", "Home"})

	var tx Transaction
	tx = ChangeEmployeeToHourly{empId, 15.00, er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

	he, ok := e.(entities.HourlyEmployee)

	if !ok {
		t.Fatalf("Failed to change employee to Hourly")
	}

	if diff := math.Abs(15.00 - he.HourlyRate); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Hourly Rate to be %f, got %f`, 15.00, he.HourlyRate)
	}

	if len(he.Timecards) != 0 {
		t.Fatalf("Should not start with any Time Cards associated")
	}
}

func TestChangeEmployeeToHourlyOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToHourly{empId, 15.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}

}

func TestChangeEmployeeToSalaried(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{empId, "Bob", "Home"})

	var tx Transaction
	tx = ChangeEmployeeToSalaried{empId, 1000.00, er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

	se, ok := e.(entities.SalariedEmployee)

	if !ok {
		t.Fatalf("Failed to change employee to Salaried")
	}

	if diff := math.Abs(1000.00 - se.Salary); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Salary to be %f, got %f`, 1000.00, se.Salary)
	}
}

func TestChangeEmployeeToSalariedOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToSalaried{empId, 1000.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}

}

func TestChangeEmployeeToCommissioned(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{empId, "Bob", "Home"})

	var tx Transaction
	tx = ChangeEmployeeToCommissioned{empId, 1000.00, 10.00, er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)

	ce, ok := e.(entities.CommissionedEmployee)

	if !ok {
		t.Fatalf("Failed to change employee to Commisssioned")
	}

	if diff := math.Abs(1000.00 - ce.Salary); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Hourly Rate to be %f, got %f`, 1000.00, ce.Salary)
	}

	if diff := math.Abs(10.00 - ce.CommissionRate); diff > 0.001 {
		t.Fatalf(`Failed to persist Employee Data properly, want Commission Rate to be %f, got %f`, 10.00, ce.CommissionRate)
	}

	if len(ce.SaleReceipts) != 0 {
		t.Fatalf("Should not start with any Sale Receipts associated")
	}
}

func TestChangeEmployeeToCommissionedOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToCommissioned{empId, 1000.00, 10.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}

}

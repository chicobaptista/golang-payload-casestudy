package usecases

import (
	"math"
	"testing"
	"time"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
)

func TestRunPayrollOneSalariedEmployeeWithNoDeductions(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	er.AddEmployee(entities.NewSalariedEmployee(empId, "Bob", "Home", 1000.00))

	tx := RunPayroll{time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), make(map[int]Paycheck), er}

	tx.Execute()

	pr := tx.Payroll

	if len(pr) != 1 {
		t.Fatalf("Failed to process Payroll, expected to have 1 entry")
	}

	pc, ok := tx.GetPaycheck(empId)
	if !ok {
		t.Fatalf(`Failed to process Payroll, expected to have paycheck for Employee %d.`, empId)
	}

	if diff := math.Abs(1000.00 - pc.Amount); diff > 0.001 {
		t.Fatalf(`Failed to process Payroll, want Paycheck Amount to be %f, got %f`, 1000.00, pc.Amount)
	}
}

func TestRunPayrollOneSalariedEmployeeOutsidePaymentDate(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	er.AddEmployee(entities.NewSalariedEmployee(empId, "Bob", "Home", 1000.00))

	tx := RunPayroll{time.Date(2023, 3, 30, 12, 30, 30, 100, time.Local), make(map[int]Paycheck), er}

	tx.Execute()

	pr := tx.Payroll

	if len(pr) != 0 {
		t.Fatalf("Failed to process Payroll, expected to have no entries")
	}

}

func TestRunPayrollOneHourlyEmployeeNoOvertimeNoDeductions(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	er.AddEmployee(entities.NewHourlyEmployee(empId, "Bob", "Home", 15.00))
	PostTimecard{empId, time.Date(2023, 3, 30, 12, 30, 30, 100, time.Local), 6.00, er}.Execute()

	tx := RunPayroll{time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), make(map[int]Paycheck), er}

	tx.Execute()

	pr := tx.Payroll

	if len(pr) != 1 {
		t.Fatalf("Failed to process Payroll, expected to have 1 entry")
	}

	pc, ok := tx.GetPaycheck(empId)
	if !ok {
		t.Fatalf(`Failed to process Payroll, expected to have paycheck for Employee %d.`, empId)
	}

	if diff := math.Abs(90.00 - pc.Amount); diff > 0.001 {
		t.Fatalf(`Failed to process Payroll, want Paycheck Amount to be %f, got %f`, 90.00, pc.Amount)
	}
}

func TestRunPayrollOneHourlyEmployeeWithOvertimeNoDeductions(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	er.AddEmployee(entities.NewHourlyEmployee(empId, "Bob", "Home", 15.00))
	PostTimecard{empId, time.Date(2023, 3, 30, 12, 30, 30, 100, time.Local), 9.00, er}.Execute()

	tx := RunPayroll{time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), make(map[int]Paycheck), er}

	tx.Execute()

	pr := tx.Payroll

	if len(pr) != 1 {
		t.Fatalf("Failed to process Payroll, expected to have 1 entry")
	}

	pc, ok := tx.GetPaycheck(empId)
	if !ok {
		t.Fatalf(`Failed to process Payroll, expected to have paycheck for Employee %d.`, empId)
	}

	if diff := math.Abs(142.50 - pc.Amount); diff > 0.001 {
		t.Fatalf(`Failed to process Payroll, want Paycheck Amount to be %f, got %f`, 142.50, pc.Amount)
	}
}

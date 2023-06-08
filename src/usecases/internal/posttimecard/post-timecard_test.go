package usecases

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"chicobaptista.github.com/payrollcasestudy/entities"
	"chicobaptista.github.com/payrollcasestudy/repositories"
	"chicobaptista.github.com/payrollcasestudy/usecases/interfaces"
)

func TestPostTimecardToHourlyEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewHourlyEmployee(empId, "Bob", "Home", 15.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostTimecard{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 6.00, er}

	tx.Execute()
	fe, _ := er.GetEmployee(empId)
	he := fe.(entities.HourlyEmployee)

	tcs := he.Timecards

	if len(tcs) != 1 {
		t.Fatalf("Should have one Timecard Recorded with the Employee")
	}

	tc := tcs[0]
	if tc.Date != time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local) {
		t.Fatalf(`Failed to persist TimeCard Data properly, want Date to be %s, got %s`, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local).String(), tc.Date.String())
	}

	if tc.Hours != 6.00 {
		t.Fatalf(`Failed to persist TimeCard Hours properly, want Date to be %.2f, got %.2f`, 6.00, tc.Hours)
	}
}

func TestPostTimecardToSalariedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewSalariedEmployee(empId, "Bob", "Home", 1000.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostTimecard{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 6.00, er}

	_, err := tx.Execute()
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is not an Hourly Employee`, empId)) {
		t.Fatalf("Should not post a Timecard to a non-Hourly Employee")
	}
}

func TestPostTimecardToCommissionedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewCommissionedEmployee(empId, "Bob", "Home", 1000.00, 10.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostTimecard{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 6.00, er}

	_, err := tx.Execute()
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is not an Hourly Employee`, empId)) {
		t.Fatalf("Should not post a Timecard to a non-Hourly Employee")
	}
}

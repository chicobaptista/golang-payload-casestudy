package usecases

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"chicobaptista.github.com/usecases/interfaces"
)

func TestPostSalesReceiptToCommissionedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewCommissionedEmployee(empId, "Bob", "Home", 1000.00, 10.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostSaleReceipt{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 100.00, er}

	tx.Execute()

	fe, _ := er.GetEmployee(empId)
	ce := fe.(entities.CommissionedEmployee)

	srs := ce.SaleReceipts

	if len(srs) != 1 {
		t.Fatalf("Should have one Sale Receipt Recorded with the Employee")
	}

	sr := srs[0]
	if sr.Date != time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local) {
		t.Fatalf(`Failed to persist Sale Receipt Date properly, want Date to be %s, got %s`, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local).String(), sr.Date.String())
	}

	if sr.Amount != 100.00 {
		t.Fatalf(`Failed to persist Sale Receipt Amount properly, want Date to be %.2f, got %.2f`, 100.00, sr.Amount)
	}
}

func TestPostSalesReceiptToHourlyEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewHourlyEmployee(empId, "Bob", "Home", 15.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostSaleReceipt{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 100.00, er}

	_, err := tx.Execute()
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is not a Commissioned Employee`, empId)) {
		t.Fatalf("Should not post a Timecard to a non-Commissioned Employee")
	}
}

func TestPostSalesReceiptToSalariedEmployee(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	e := entities.NewSalariedEmployee(empId, "Bob", "Home", 1000.00)
	er.AddEmployee(e)

	var tx interfaces.Transaction
	tx = PostSaleReceipt{empId, time.Date(2023, 3, 31, 12, 30, 30, 100, time.Local), 100.00, er}

	_, err := tx.Execute()
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is not a Commissioned Employee`, empId)) {
		t.Fatalf("Should not post a Timecard to a non-Commissioned Employee")
	}
}

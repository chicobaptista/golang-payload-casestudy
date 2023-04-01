package usecases

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
)

func TestChangeEmployeeName(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.HoldingPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

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
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.HoldingPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

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
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.HoldingPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

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
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.HoldingPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

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
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.HoldingPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

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

func TestChangeEmployeeToHoldingPaymentMethod(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

	var tx Transaction
	tx = ChangeEmployeeToHolding{empId, er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)
	be, _ := e.(entities.BaseEmployee)

	_, ok := be.PaymentMethod.(entities.HoldingPaymentMethod)

	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee Payment method to be of type Holding")
	}
}

func TestChangeEmployeeToHoldingPaymentMethodOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToHolding{empId, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}
}

func TestChangeEmployeeToMailPaymentMethod(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

	var tx Transaction
	tx = ChangeEmployeeToMail{empId, "Work", er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)
	be, _ := e.(entities.BaseEmployee)

	pm, ok := be.PaymentMethod.(entities.MailPaymentMethod)

	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee Payment method to be of type Mail")
	}

	if pm.Address != "Work" {
		t.Fatalf(`Failed to persist Employee Data properly, want Payment Mailing Address to be %s, got %s`, "Work", pm.Address)

	}
}

func TestChangeEmployeeToMailPaymentMethodOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToMail{empId, "Work", er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}
}

func TestChangeEmployeeToDirectPaymentMethod(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

	var tx Transaction
	tx = ChangeEmployeeToDirect{empId, "Agency", "Account", er}

	tx.Execute()

	e, _ := er.GetEmployee(empId)
	be, _ := e.(entities.BaseEmployee)

	pm, ok := be.PaymentMethod.(entities.DirectPaymentMethod)

	if !ok {
		t.Fatalf("Failed to persist Employee Data properly, want Employee Payment method to be of type Mail")
	}

	if pm.Agency != "Agency" {
		t.Fatalf(`Failed to persist Employee Data properly, want Payment Agency to be %s, got %s`, "Agency", pm.Agency)

	}
	if pm.Account != "Account" {
		t.Fatalf(`Failed to persist Employee Data properly, want Payment Account to be %s, got %s`, "Account", pm.Account)

	}
}

func TestChangeEmployeeToDirectPaymentMethodOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeToDirect{empId, "Agency", "Account", er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Category of a Non-Existing Employee.")
	}
}

func TestChangeEmployeeToUnionMember(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})
	unionId := 86

	var tx Transaction
	tx = ChangeEmployeeToMember{empId, unionId, 100.00, er}

	tx.Execute()

	unionEmp, _ := er.GetUnionMember(unionId)
	if len(unionEmp.Charges) != 0 {
		t.Fatalf("Should not have any Union Charges Recorded with the Union Member")
	}
	if diff := math.Abs(100.00 - unionEmp.Dues); diff > 0.001 {
		t.Fatalf(`Failed to persist Union Member properly, want Dues to be %.2f, got %.2f`, 100.00, unionEmp.Dues)
	}

}

func TestChangeEmployeeToUnionMemberOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	unionId := 86

	var tx Transaction
	tx = ChangeEmployeeToMember{empId, unionId, 100.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Affiliation of a Non-Existing Employee.")
	}
}

func TestChangeEmployeeToUnionMemberOfAlreadyMember(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	unionId := 86
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.UnionAffiliation{Id: unionId}})

	var tx Transaction
	tx = ChangeEmployeeToMember{empId, unionId, 100.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is already a Member`, empId)) {
		t.Fatalf("Should not Change the Affiliation of an Employee that is already affiliated.")
	}
}

func TestChangeEmployeeToUnionMemberOfExistingMemberId(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	unionId := 86
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})
	er.PutUnionMember(entities.UnionMember{Id: unionId, Dues: 10.00, Charges: make([]entities.UnionCharge, 0)})

	var tx Transaction
	tx = ChangeEmployeeToMember{empId, unionId, 100.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Member %d is already registered`, unionId)) {
		t.Fatalf("Should not Affiliate an Employee with the Id of an existing Member")
	}
}

func TestChangeEmployeeRemoveUnionMember(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	unionId := 86
	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.UnionAffiliation{Id: unionId}})
	er.PutUnionMember(entities.UnionMember{Id: unionId, Dues: 10.00, Charges: make([]entities.UnionCharge, 0)})

	var tx Transaction
	tx = ChangeEmployeeRemoveMember{empId, er}

	tx.Execute()
	e, _ := er.GetEmployee(empId)
	be, _ := e.(entities.BaseEmployee)

	_, ok := be.Affiliation.(entities.NullAffiliation)
	if !ok {
		t.Fatalf(`Failed to persist Employee Data properly, want Affiliation to be Null`)
	}
}

func TestChangeEmployeeRemoveUnionMemberOfNonExisting(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = ChangeEmployeeRemoveMember{empId, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not Change the Affiliation of a Non-Existing Employee.")
	}
}

func TestChangeEmployeeRemoveUnionMemberOfNonMember(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	er.AddEmployee(entities.BaseEmployee{Id: empId, Name: "Bob", Address: "Home", PaymentMethod: entities.MailPaymentMethod{}, Affiliation: entities.NullAffiliation{}})

	var tx Transaction
	tx = ChangeEmployeeRemoveMember{empId, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d is not a Member`, empId)) {
		t.Fatalf("Should not Change the Affiliation of an Employee that is already affiliated.")
	}
}

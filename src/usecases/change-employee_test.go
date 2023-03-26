package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"fmt"
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

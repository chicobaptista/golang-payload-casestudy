package usecases

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"fmt"
	"strings"
	"testing"
)

func TestDeleteExistingEmployee(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	er.AddEmployee(entities.NewSalariedEmployee(empId, "Bob", "Home", 1000.00))

	var tx Transaction
	tx = DeleteEmployee{empId, er}
	tx.Execute()

	e, _ := er.GetEmployee(empId)

	if e != nil {
		t.Fatalf("Failed to delete Employee Data properly")

	}
}

func TestDeleteNonExistingEmployee(t *testing.T) {
	var er EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	empId := 1

	var tx Transaction
	tx = DeleteEmployee{empId, er}
	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Employee %d not found`, empId)) {
		t.Fatalf("Should not delete a Non Existing Employee")
	}
}

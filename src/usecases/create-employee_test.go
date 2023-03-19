package usecases

import (
	"chicobaptista.github.com/repositories"
	"testing"
)

func TestAddSalariedEmployee(t *testing.T) {
	er := repositories.MakeInMemoryEmployeeRepository()

	empId := 1
	var tx Transaction
	tx = AddSalariedEmployee{empId, "Bob", "Home", 1000.00, er}
	tx.Execute()
	e := er.GetEmployee(empId)
	if e.Name != "Bob" {
		t.Fatalf(`Failed to persist Employee Data properly, want Name to be %q, got %v`, "Bob", e.Name)
	}

}

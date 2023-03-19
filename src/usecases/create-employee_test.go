package usecases

import (
	"testing"
)

func TestAddSalariedEmployee(t *testing.T) {
	empId := 1
	var tx Transaction
	tx = AddSalariedEmployee{empId, "Bob", "Home", 1000.00}
	tx.Execute()

}

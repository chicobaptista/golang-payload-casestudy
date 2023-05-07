package usecases

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/repositories"
	"chicobaptista.github.com/usecases/interfaces"
)

func TestPostUnionchargeToUnionMember(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	unionId := 1
	er.PutUnionMember(entities.UnionMember{Id: unionId, Dues: 20.00, Charges: make([]entities.UnionCharge, 0)})

	var tx interfaces.Transaction
	tx = PostUnionCharge{unionId, 10.00, er}

	tx.Execute()

	unionEmp, _ := er.GetUnionMember(unionId)
	if len(unionEmp.Charges) != 1 {
		t.Fatalf("Should have one Union Charge Recorded with the Union Member")
	}

	uc := unionEmp.Charges[0]
	if diff := math.Abs(10.00 - uc.Amount); diff > 0.001 {
		t.Fatalf(`Failed to persist Union Charge properly, want Amount to be %.2f, got %.2f`, 10.00, uc.Amount)
	}
}

func TestPostUnionchargeToNonMember(t *testing.T) {
	var er interfaces.EmployeeRepository
	er = repositories.MakeInMemoryEmployeeRepository()

	unionId := 1

	var tx interfaces.Transaction
	tx = PostUnionCharge{unionId, 10.00, er}

	_, err := tx.Execute()

	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf(`Member %d not found.`, unionId)) {
		t.Fatalf("Should not be able to Post a Union Charge to a non-Member")

	}
}

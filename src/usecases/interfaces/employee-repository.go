package interfaces

import "chicobaptista.github.com/payrollcasestudy/entities"

type EmployeeRepository interface {
	GetEmployee(empId int) (entities.Employee, bool)
	AddEmployee(e entities.Employee)
	DeleteEmployee(empId int)
	GetUnionMember(memberId int) (entities.UnionMember, bool)
	PutUnionMember(um entities.UnionMember)
	DeleteUnionMember(unionId int)
}

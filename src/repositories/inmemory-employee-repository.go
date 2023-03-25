package repositories

import "chicobaptista.github.com/entities"

type InMemoryEmployeeRepository struct {
	employees    map[int]entities.Employee
	unionMembers map[int]entities.UnionMember
}

func (er InMemoryEmployeeRepository) GetEmployee(empId int) (entities.Employee, bool) {
	e, ok := er.employees[empId]
	return e, ok
}

func (er InMemoryEmployeeRepository) AddEmployee(e entities.Employee) {
	er.employees[e.GetId()] = e
}

func (er InMemoryEmployeeRepository) DeleteEmployee(empId int) {
	delete(er.employees, empId)
}

func (er InMemoryEmployeeRepository) GetUnionMember(memberId int) (entities.UnionMember, bool) {
	m, ok := er.unionMembers[memberId]
	return m, ok
}

func (er InMemoryEmployeeRepository) PutUnionMember(um entities.UnionMember) {
	er.unionMembers[um.Id] = um
}

func MakeInMemoryEmployeeRepository() InMemoryEmployeeRepository {
	return InMemoryEmployeeRepository{make(map[int]entities.Employee), make(map[int]entities.UnionMember)}
}

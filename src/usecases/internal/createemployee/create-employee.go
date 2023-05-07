package createemployee

import (
	"chicobaptista.github.com/entities"
	"chicobaptista.github.com/usecases/interfaces"
)

var HOURLY = "H"
var SALARY = "S"
var COMISSIONED = "C"

type CreateEmployeeRequest struct {
	Type          string
	Id            int
	Name          string
	Address       string
	HourlyRate    float64
	Salary        float64
	ComissionRate float64
}

func MakeCreateEmployee(req CreateEmployeeRequest, er interfaces.EmployeeRepository) CreateEmployee {

	switch req.Type {
	case HOURLY:
		return CreateEmployee{CreateHourlyEmployee{req.Id, req.Name, req.Address, req.HourlyRate}, er}
	case SALARY:
		return CreateEmployee{CreateSalariedEmployee{req.Id, req.Name, req.Address, req.Salary}, er}
	case COMISSIONED:
		return CreateEmployee{CreateCommissionedEmployee{req.Id, req.Name, req.Address, req.Salary, req.ComissionRate}, er}
	default:
		panic("Employee type not supported")
	}
}

type CreateEmployee struct {
	ce    CreateEmployeeBehavior
	eRepo interfaces.EmployeeRepository
}

func (tx CreateEmployee) Execute() (success bool, err error) {
	e := tx.ce.generateEmployee()
	tx.saveEmployee(e)
	return true, nil
}

func (tx CreateEmployee) saveEmployee(e entities.Employee) {
	tx.eRepo.AddEmployee(e)
}

type CreateEmployeeBehavior interface {
	generateEmployee() entities.Employee
}

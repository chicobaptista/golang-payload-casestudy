# Create Employee Use Case

> A new employee is added by the receipt of an `AddEmp` transaction. This transaction contains the employee's name, address, and assigned employee number. The transaction has the following three forms:
>
> ```xml
> AddEmp <EmpID> "<name>" "<address>" H <hourly-rate>
> AddEmp <EmpID> "<name>" "<address>" S <monthly-salary>
> AddEmp <EmpID> "<name>" "<address>" C <monthly-salary> > <commission-rate>
> ```
>
> The employee record is created with its fields assigned appropriately.
>
> - **Alternative:**
>
>   **An error in the transaction structure**
>
>   If the transaction structure is innapropriate, it is printed out in an error message, and no action is taken.

## Strategy

As all forms of this use case follow the same basic steps, a template method pattern is used to orchestrate the execution and concrete implementations of each case deal with the specifics of how to generate a Concrete Employee entity.

## Class Diagram

```plantuml
@startuml
skinparam backgroundColor #EEEBDC
interface Transaction {
    Execute()
}
class CreateEmployee{
    saveEmployee(Employee)
}
CreateEmployee--|>Transaction

interface EmployeeRepository
CreateEmployee-->EmployeeRepository:persists

interface CreateEmployeeBehavior{
    generateEmployee() Employee
}
CreateEmployee*-->CreateEmployeeBehavior:delegates generation

class CreateSalariedEmployee
CreateSalariedEmployee--|>CreateEmployeeBehavior
CreateSalariedEmployee-->SalariedEmployee
class CreateComissionedEmployee
CreateComissionedEmployee--|>CreateEmployeeBehavior
CreateComissionedEmployee-->ComissionedEmployee
class CreateHourlyEmployee
CreateHourlyEmployee--|>CreateEmployeeBehavior
CreateHourlyEmployee-->HourlyEmployee

interface Employee
class HourlyEmployee
HourlyEmployee--|>Employee
class SalariedEmployee
SalariedEmployee--|>Employee
class ComissionedEmployee
ComissionedEmployee--|>Employee
CreateEmployeeBehavior-->Employee:generates


@enduml
```

## Sequence Diagram

```plantuml
@startuml
    skinparam backgroundColor #EEEBDC
    boundary CreateEmpController
    control CreateEmployee
    CreateEmpController->CreateEmployee: empID, name, address, empType, ...
    CreateEmpController->CreateEmployee: "execute()"
    activate CreateEmployee
    entity Employee

    CreateEmployee->Employee: "NewConcreteEmployee()"
    activate Employee
    Employee-->CreateEmployee:ConcreteEmp
    deactivate Employee
    database EmployeeRepository
    CreateEmployee->EmployeeRepository: "AddEmployee(ConcreteEmp)"
    activate EmployeeRepository
    EmployeeRepository-->CreateEmployee: ok
    deactivate EmployeeRepository
    CreateEmployee-->CreateEmpController:ok
    deactivate CreateEmployee
@enduml
```

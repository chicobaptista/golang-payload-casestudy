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

## Sequence Diagram

```plantuml
@startuml
    skinparam backgroundColor #EEEBDC
    boundary CreateEmpController
    control CreateEmployee
    CreateEmpController->CreateEmployee: empID, name, address, empType, ...
    CreateEmployee->CreateEmployee: "execute()"
    entity Employee
    CreateEmployee->Employee: "NewConcreteEmployee()"
    Employee->CreateEmployee:ConcreteEmp
    database EmployeeRepository
    CreateEmployee->EmployeeRepository: "AddEmployee(ConcreteEmp)"
    EmployeeRepository->CreateEmployee:ok
    CreateEmployee->CreateEmpController:ok
@enduml
```

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
interface CreateEmployeeBehavior{
    generateEmployee() Employee
}
class CreateHourlyEmployee
CreateHourlyEmployee--|>CreateEmployeeBehavior
CreateSalariedEmployee--|>CreateEmployeeBehavior
CreateComissionedEmployee--|>CreateEmployeeBehavior
CreateEmployee*-->CreateEmployeeBehavior:delegates generation
interface Employee
class HourlyEmployee
HourlyEmployee--|>Employee
class SalariedEmployee
SalariedEmployee--|>Employee
class ComissionedEmployee
ComissionedEmployee--|>Employee
CreateEmployeeBehavior-->Employee:generates
interface EmployeeRepository
CreateEmployee-->EmployeeRepository:persists
@enduml
```

# Payroll System Case Study

This is an implementation of the Payroll System Case Study presented by Uncle Bob on the **Agile Software Development: Principles, Practices and Patterns** book, using Golang.

The purpose is to develop a Payroll System that can calculate and process the payment of several employees, according to different rules regarding contract types, rates, commission, etc.

## Specification of the Payroll System

This system consists of a database of the employees in the company and their associated data, such as time cards. The system must pay each employee. Employees must be paid the correct amount, on time, by the method that they specify. Also, various deductions must be taken from their pay.

- Some employees work by the hour. They are paid an hourly rate that is one of the fields in their employee record. They submit daily time cards that record the date and the umber of hours worked. If they work more than 8 hours per day, they are paid 1.5 times their normal rate for those extra hours. They are paid every Friday.
- Some employees are paid a flat salary. They are paid on the last working day of the month. Their monthly salary is one of the fields in their employee record.
- Some of the salaried employees are also paid a commission based on their sales. They submit sales receipts that record the date and the amount of the sale. Their commission rate is a field in their employee record. They are paid every other Friday.
- Employees can select their method of payment. They may have thair paychecks mailed to the postal address of their choice: they may have their paychecks held for pickuop by the Paymaster; or they can request that their paychecks be directly deposited into the bank account of their choice.
- Some employees belong to the union. Their employee record has a field for the weekly dues rate. Their dues must be deducted from their pay. Also, the union may assess service charges against individual union members from time to time. These service charges are submitted by the union on a weekly basis and must be deducted from the appropriate employee's next pay amount.
- The payroll application will run once each working day and pay the appropriate employees on that day. The system will be told to what date the employees are to be paid, so it will calculate payments from the last time the employee was paid up to the specified date.

## Use Cases

The following use cases must be implemented in the system:

### Use Case 1: Add New Employee

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

### Use Case 2: Deleting an Employee

> Employees are deleted when a `DelEmp` transaction is received. The form of the trsnsaction id as follows:
>
> ```xml
> DelEmp <EmpID>
> ```
>
> When this transaction is received, the appropriate employee record is deleted.
>
> - **Alternative**
>
>   **Invalid or unknown `EmpID`**
>
>   If the `<EmpID>` field is not strucutred correclty, ot if it does not refer to a valid employee record, then the transaction is printed with an error message, and no other action is taken.

### Use Case 3: Posting a `Time Card`

> Upon receipt os a `TimeCard` transaction, the ssytem will > create a time-card record and associate it with the appropriate employee record.
>
> ```xml
> TimeCard <EmpId> <date> <hours>
> ```
>
> - **Alternative 1:**
>
>   **The selected employee is not hourly**
>
>   The system will print an appropriate error message and > take no further action.
>
> - **Alternative 2:**
>
>   **An error in the transaction structure**
>
>   The system will print an appropriate error message and > take no further action.

### Use Case 4: Posting a `Sales Receipt`

> Upon receipt of the `SalesReceipt` transaction, the system will create a new sales-receipt record and associate it with the approprate commissioned employee.
>
> ```xml
> SalesReceipt <EmpId> <date> <amount>
> ```
>
> - **Alternative 1:**
>
>   **The selected employee is not commissioned**
>
>   The system will print an appropriate error message and take no further action.
>
> - **Alternative 2:**
>
>   **An error in the transaction structure**
>
>   The system will print an approprate error message and take no further action.

### Use Case 5: Posting a Union Service Charge

> Upon receipt of this transaction, the system will create a service-charge record and associate it with the appropriate union member.
>
> ```xml
> ServiceCharge <memberID> <amount>
> ```
>
> - **Alternative:**
>
>   **Poorly formated transaction**
>
>   If the transaction is not well formed or if the `<memberID>` does not refer to an existing union member, then the transaction is printed with an approppriate error message and no further action is taken.

### Use Case 6: Changing Employee Details

> Upon receipt os this transaction, the system will alter one of the details of the appropriate employee recors. There are several possible variations to this transaction.
>
> ```xml
> ChgEmp <EmpID> Name <name>                      Change Employee Name
> ChgEmp <EmpID> Address <address>                Change Employee Address
> ChgEmp <EmpID> Hourly <hourly-rate>             Change to Hourly
> ChgEmp <EmpID> Salaried <salary>                Change to Salaried
> ChgEmp <EmpID> Commissioned <salary> <rate>     Change to Commissioned
> ChgEmp <EmpID> Hold                             Hold Paycheck
> ChgEmp <EmpID> Direct <bank> <account>          Direct Deposit
> ChgEmp <EmpID> Mail <address>                   Mail Paycheck
> ChgEmp <EmpID> Member <memberID> Dues <rate>    Put Employee in Union
> ChgEmp <EmpID> NoMember                         Remove Employee from Union
> ```
>
> - **Alternative:**
>
>   **Transaction Errors**
>
>   If the structure of the transaction is improper or `<EmpID>`does not refer to a real employee or `<memberID>`already refers to a member, then print a suitable error and take no further action.

### Use Case 7: Run the Payroll for Today

> Upon receipt os the Payday transaction, the system finds all those employees that should be paid upon the specified date. The system then determines how much they are owed and pays them according to their selected payment method.
>
> ```xml
> Payday <date>
> ```

package Employees

import "fmt"

// Employee interface
type Employee interface {
	GetDetails() string
}

// FullTimeEmployee struct
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTimeEmployee[ID=%d, Name=%s, Salary=%d]", fte.ID, fte.Name, fte.Salary)
}

// PartTimeEmployee struct
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (pte PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("PartTimeEmployee[ID=%d, Name=%s, HourlyRate=%d, HoursWorked=%.2f, TotalPay=%.2f]",
		pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked, float64(pte.HourlyRate)*float64(pte.HoursWorked))
}

// Company struct
type Company struct {
	employees map[string]Employee
}

// NewCompany constructor
func NewCompany() *Company {
	return &Company{
		employees: make(map[string]Employee),
	}
}

// AddEmployee adds an employee
func (c *Company) AddEmployee(empID string, emp Employee) {
	c.employees[empID] = emp
	fmt.Println("Employee added:", empID)
}

// ListEmployees prints all employees
func (c *Company) ListEmployees() {
	for id, emp := range c.employees {
		fmt.Println("ID:", id, emp.GetDetails())
	}
}

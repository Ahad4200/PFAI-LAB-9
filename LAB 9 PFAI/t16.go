package main

import "fmt"

// Employee struct with Name, Position, Salary
type Employee struct {
	Name     string
	Position string
	Salary   float64
}

// highestPaid returns the employee with the highest salary
func highestPaid(employees []Employee) Employee {
	if len(employees) == 0 {
		return Employee{}
	}

	highest := employees[0]
	for _, emp := range employees {
		if emp.Salary > highest.Salary {
			highest = emp
		}
	}

	return highest
}

func main() {
	// Create a slice of 5 employees
	employees := []Employee{
		{Name: "Alice Johnson", Position: "Software Engineer", Salary: 85000},
		{Name: "Bob Smith", Position: "Project Manager", Salary: 92000},
		{Name: "Carol Davis", Position: "CTO", Salary: 120000},
		{Name: "David Wilson", Position: "UI Designer", Salary: 78000},
		{Name: "Eva Brown", Position: "DevOps Engineer", Salary: 90000},
	}

	// Find and print the highest paid employee
	topEarner := highestPaid(employees)
	fmt.Printf("Highest paid employee:\n")
	fmt.Printf("Name: %s\n", topEarner.Name)
	fmt.Printf("Position: %s\n", topEarner.Position)
	fmt.Printf("Salary: $%.2f\n", topEarner.Salary)
}

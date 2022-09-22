package blockchain_tasks

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}
type company struct {
	companyName string
	employees   []employee
}

func AddCompany(employees []employee) {
	companyy := company{"BigAnalytics", employees[:]}
	PrintCompany(companyy)
}
func AddEmployees() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Amina", 90000, "Front-End Developer"}
	emp3 := employee{"Alina", 100000, "Project Manager"}

	employees := [3]employee{emp1, emp2, emp3}
	AddCompany(employees[:])

}
func PrintCompany(companyy company) {
	fmt.Println("Company: ", companyy.companyName)
	for i := 0; i < 3; i++ {
		fmt.Println("\nEmployee # ", i+1)
		fmt.Println("Name: ", companyy.employees[i].name)
		fmt.Println("Salary: ", companyy.employees[i].salary)
		fmt.Println("Job Description: ", companyy.employees[i].position)
	}
}

// Call AddEmployees in main for all functionality

package blockchain_tasks

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func StructTesting(person1, person2 Person) {
	fmt.Println("Name of 1st person is: ", person1.name)
	fmt.Println("Age of 1st person is: ", person1.age)
	fmt.Println("Job of 1st person is: ", person1.job)
	fmt.Println("Salary of 1st person is: ", person1.salary)
	fmt.Println("\n\nName of 2nd person is: ", person2.name)
	fmt.Println("Age of 2nd person is: ", person2.age)
	fmt.Println("Job of 2ne person is: ", person2.job)
	fmt.Println("Salary of 2nd person is: ", person2.salary)

}

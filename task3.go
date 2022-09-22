package blockchain_tasks

import (
	"fmt"
	"strings"
)

type Student struct {
	rollNumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollNumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}
func Print(student *StudentList) {
	for i := 0; i < len(student.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Println("Age: ", student.list[i].rollNumber)
		fmt.Println("Name: ", student.list[i].name)
		fmt.Println("Address: ", student.list[i].address)
	}

}

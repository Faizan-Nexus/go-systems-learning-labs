package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	s1 := Student{"Zain", 01, "CSE"}
	s2 := Student{"Younis", 02, "CSE"}
	s3 := Student{"Waqar", 32, "CSE"}
	fmt.Printf("Student Details are %+v , %+v, %+v\n", s1, s2, s3)
	fmt.Println()
	s1.GetID("Zain")
	s1.ChangeDepartmnet(1, "CS")
	fmt.Println("Zain", s1)
}

type Student struct {
	Name       string
	ID         int
	Department string
}

func (s Student) GetID(name string) {
	if s.Name == name {
		fmt.Printf("ID of %s is %d\n", name, s.ID)
	} else {
		fmt.Printf("No student found with name %s\n", name)
	}
}

func (s *Student) ChangeDepartmnet(id int, depart string) {
	if s.ID == id {
		s.Department = depart
		fmt.Println("New Department is:", s.Department)
	} else {
		fmt.Printf("No student found with id %v\n", id)
	}
}

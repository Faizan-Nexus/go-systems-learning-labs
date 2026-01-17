package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruits = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Fruits slice:", fruits)
	fmt.Printf("Type is: %T\n", fruits)

	fruits = append(fruits, "Mango")
	fmt.Println("Updated fruit slices is: ", fruits)

	// Another way of declaration
	grades := make([]int, 4)
	grades[0] = 78
	grades[1] = 87
	grades[2] = 83
	grades[3] = 84

	fmt.Println("The grades are", grades)

	//some methods
	sort.Ints(grades)
	fmt.Println("The sorted slices are as ", grades)
	fmt.Println("Is slice grade is sorted:", sort.IntsAreSorted(grades))

	//some more
	var courses = []string{"Java", "AI ML", "DSD", "DLD"}
	courses = append(courses[:2], courses[3:]...)
	fmt.Println("Courses are optomized for memory reallocation", courses)
}

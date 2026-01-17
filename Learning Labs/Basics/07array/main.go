package main

import "fmt"

func main() {
	fmt.Println("This is the main function in the 07array package")

	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fruits[3] = "Date"

	fmt.Println("Fruits array:", fruits)
	fmt.Println("Length of fruits array:", len(fruits))

	var vegetables = [4]string{"Carrot", "Broccoli", "Spinach", "Potato"}
	fmt.Println("Vegetables array:", vegetables)

	var names = [4]string{"Faizan", "Abdul Rehman", "Zohiab"}
	fmt.Println("Names includes",names)
	fmt.Printf("names: %v\n", names)
}

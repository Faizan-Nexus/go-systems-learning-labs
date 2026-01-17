package main

//Pointers are used to refer to memory location of variables

import "fmt"

func main() {
	fmt.Println("Go Pointers")

	// var ptr *int
	// fmt.Println("The default value of a pointer is: ", ptr)

	num := 07
	ptr := &num
	fmt.Println("The pointer is refrencing to", ptr, "memory address")
	fmt.Println("The value at that address is: ", *ptr)

	fmt.Println("\nWe can perform operations on pointers as well")
	*ptr = *ptr + 2
	fmt.Println("After performing addition the value at that address is now: ", *ptr)

	fmt.Println("\nWe can also create a pointer to a pointer")
	var ptr2 = &ptr
	fmt.Println("The value of ptr2 is: ", ptr2)
	fmt.Println("The address it is pointing to: ", *ptr2)
	fmt.Println("The value at that address is: ", **ptr2)
}

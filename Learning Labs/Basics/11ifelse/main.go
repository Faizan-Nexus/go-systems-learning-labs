package main

import "fmt"

func main() {
	fmt.Println("Conditional if else in Go")
	var result string

	age := 23
	if age > 18 {
		result = "You are adult"
	} else if age > 21 {
		result = "You are old enough"
	} else {
		result = "You are still a kid"
	}
	fmt.Println(result)

	//Another Synatax
	if num := 5; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

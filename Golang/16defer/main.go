package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Defer in Golang")

	defer fmt.Println("Happy New Year!")
	for i := 1; i < 6; i++ {
		defer fmt.Println(i)
	}
}


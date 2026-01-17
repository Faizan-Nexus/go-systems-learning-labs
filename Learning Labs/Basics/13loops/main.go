package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	var fruits = []string{"Banana", "Apple", "Mango", "Orange"}

	for i, j := range fruits {
		fmt.Printf("Value at %v is %v \n", i+1, j)
	}
	num := 0
	for num <= 10 {

		// if num == 2 {
		// 	goto qml
		// }

		if num == 5 {
			num++
			fmt.Println("We just wana skip", num)
			continue
		}

		if num == 10 {
			fmt.Println("We are breaking up")
			break
		}
		fmt.Println("Counting", num)
		num++
	}

// qml:
// 	fmt.Println("QML is shortcut for Quantum Machine Learning")
}

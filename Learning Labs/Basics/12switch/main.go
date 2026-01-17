package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1 and you can move 1 step")
	case 2:
		fmt.Println("Value is 2 and you can move 2 steps")
	case 3:
		fmt.Println("Value is 3 and you can move 3 steps")
	case 4:
		fmt.Println("Value is 4 and you can move 4 steps")
	case 5:
		fmt.Println("Value is 5 and you can move 5 steps")
	case 6:
		fmt.Println("Value is 6 and you can move 6 steps")
		fallthrough
	default:
		fmt.Println("Wow you got a Sixer!")
	}
}

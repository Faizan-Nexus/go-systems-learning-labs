package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")

	// Comma error or Try Catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is", input)
	fmt.Printf("The type of input is %T ", input)

}

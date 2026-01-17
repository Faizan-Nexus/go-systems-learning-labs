package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Type Conversion")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the Rating for movie")
	input, _ := reader.ReadString('\n')

	fmt.Println("You rated movie", input, "stars")

	fmt.Printf("The type of user input is %T \n", input)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After conversion and adding 1", num+1)
		fmt.Printf("Your input type is %T \n", num)
	}

	read := bufio.NewReader(os.Stdin)

	user, _ := read.ReadString('\n')

	fmt.Println(user)
}

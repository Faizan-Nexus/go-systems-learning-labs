package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cal() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number")
	str1, _ := reader.ReadString('\n')
	num1, _ := strconv.ParseFloat(strings.TrimSpace(str1), 64)
	fmt.Println("Enter Operation to perform (+,-,*,/,%):")
	str2, _ := reader.ReadString('\n')
	oper := strings.TrimSpace(str2)

	fmt.Println("Enter 2nd number number")
	str3, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseFloat(strings.TrimSpace(str3), 64)

	var operators = []string{"+", "-", "*", "/", "%"}

	for i := range operators {
		if oper == operators[i] {
			break
		} else if i == 4 {
			fmt.Println("Invalid Operator!")
		}
	}

	switch oper {
	case "+":
		fmt.Printf("%v + %v = %v", num1, num2, (num1 + num2))
	case "-":
		fmt.Printf("%v - %v = %v", num1, num2, (num1 - num2))
	case "*":
		fmt.Printf("%v * %v = %v", num1, num2, (num1 * num2))
	case "/":
		fmt.Printf("%v / %v = %v", num1, num2, (num1 / num2))
	case "%":
		fmt.Printf("%v %% %v = %v", int(num1), int(num2), int(num1)%int(num2))
	default:
		Cal()
	}
}

func CalC() {
	reader := bufio.NewReader(os.Stdin)

	var num1, num2 float64

	fmt.Println("Enter 1st Number")
	fmt.Scanf("%f", &num1)

	var oper string
	for {
		fmt.Println("Enter operation (+, -, *, /, %):")
		opStr, _ := reader.ReadString('\n')
		oper = strings.TrimSpace(opStr)
		if oper == "+" || oper == "-" || oper == "*" || oper == "/" || oper == "%" {
			break
		} else {
			fmt.Println("Invalid operator. Try again.")
		}
	}

	fmt.Println("Enter 2nd Number")
	fmt.Scanf("%f", &num2)

	switch oper {
	case "+":
		fmt.Printf("%v + %v = %v\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%v - %v = %v\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%v * %v = %v\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		fmt.Printf("%v / %v = %v\n", num1, num2, num1/num2)
	case "%":
		if int(num2) == 0 {
			fmt.Println("Error: Modulo by zero")
			return
		}
		fmt.Printf("%v %% %v = %v\n", int(num1), int(num2), int(num1)%int(num2))
	default:
		fmt.Println("Unknown error occurred")
	}
}

func main() {
	fmt.Println("Welcome to functions")
	CalC()
}

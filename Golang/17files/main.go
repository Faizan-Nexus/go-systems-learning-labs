package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golnag")
	intro := "Hello my name is Muhammad Faizan Anjum Shah." +
		"I am a Computer Systems Engineering Student"
	file, err := os.Create("./myintro.txt")

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	io.WriteString(file, intro)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Length of file is", length)
	defer file.Close()
	readFile("./myintro.txt")

}

func readFile(file string) {
	databytes, err := ioutil.ReadFile(file)

	// if err != nil {
	// 	panic(err)
	// }
	tryCatch(err)

	fmt.Println("Text inside the file:", databytes)

	fmt.Println("Text inside the file:", string(databytes))
}

func tryCatch(err error) {
	if err != nil {
		panic(err)
	}
}

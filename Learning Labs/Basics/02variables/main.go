package main

import "fmt"

// outside := "Out of the World" //This is not allowed

var oustide = "This is Allowed"

var Login bool = true

func main() {
	var username string = "Faizan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T  \n", username)

	//Alliases
	var hero = "Tom Cruise"
	fmt.Println(hero)

	villan := "Batman"
	fmt.Println("Ispider man ispider man :", villan)

	fmt.Println("Variable declared ouside", oustide)

	fmt.Println("This special type of variable has public access as it first letter is capital", Login)
}

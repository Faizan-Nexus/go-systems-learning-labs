package main

import "fmt"

func main() {
	fmt.Println("Structs in GOlang")
	// NO inhertance or no super or parent

	user1 := User{"Faizan", "fantasticfaizan@go.dev", 21, true}
	fmt.Println(user1)
	fmt.Printf("The user details are %+v \n", user1)
	fmt.Println()
	fmt.Printf("The user %v has email %v", user1.Name, user1.Email)
}

type User struct {
	Name     string
	Email    string
	Age      int
	IsOnline bool
}

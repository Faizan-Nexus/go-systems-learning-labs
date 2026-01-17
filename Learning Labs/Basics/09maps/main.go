package main

import "fmt"

func main() {
	fmt.Println(" Maps in golang")

	langs := make(map[string]string) //first key then value

	langs["Go"] = "Golang"
	langs["Py"] = "Python"
	langs["Js"] = "Java Script"
	langs["C"] = "Compiler"

	fmt.Println("List of all languages", langs)
	fmt.Println("The Js key has", langs["Js"], "Programming langugae")
	fmt.Println("Length of our map is: ", len(langs))

	delete(langs, "C")
	fmt.Println("List of updated langugaes", langs)

	for key, val := range langs {
		fmt.Printf("The key %v has value: %v \n", key, val)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go")
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)

	fmt.Println("This time looks too messy let's fomrat")
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2004, time.May, 19, 19, 23, 21, 0, time.UTC)
	fmt.Println("My Date of Birth is: ", createDate)

	// You can build this program to be able to run on other paltforms as well
	//using the follwing command $env:GOOS = "darwin" go build
}

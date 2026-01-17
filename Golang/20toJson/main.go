package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string
	Age        int
	Id         int
	Department []string `json:"dept sem"`
}

func main() {
	fmt.Println("Welcome to JSON converiosn")
	fmt.Println()
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	studentData := []course{
		{"Zain", 24, 0001, []string{"CSE", "VI"}},
		{"Zohaib", 21, 0035, []string{"CSE", "VI"}},
		{"Qurban", 22, 0042, []string{"CS", "VI"}},
		{"Mumtaz", 21, 0013, nil},
	}
	// final, err1 := json.Marshal(studentData)
	final, err1 := json.MarshalIndent(studentData, "", "\t")
	TryCatch(err1)

	fmt.Println("Values are of Json as follows", string(final))
	fmt.Printf("Type of final is %T:", final)

}

func TryCatch(err error) {
	if err != nil {
		panic(err)
	}
}

func DecodeJson() {
	//Take some data from web in json format
	data := []byte(`[
	{"Name": "Zain", "Age": 24, "Id": 1, "Department": ["CSE", "VI"]},
	{"Name": "Zohaib", "Age": 21, "Id": 35, "Department": ["CSE", "VI"]},
	{"Name": "Qurban", "Age": 22, "Id": 42, "Department": ["CS", "VI"]}
	]`)

	var todos []course
	// Decode the JSON data into the todos slice

	err := json.Unmarshal(data, &todos)
	TryCatch(err)

	fmt.Println("Decoded JSON data from web:")
	for key, todo := range todos {
		fmt.Println("Index:", key)
		fmt.Println("Name:", todo.Name)
		fmt.Println("Age:", todo.Age)
		fmt.Println("Id:", todo.Id)
		fmt.Println("Department:", todo.Department)
		fmt.Println()
	}
}

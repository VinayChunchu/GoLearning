package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name     string `json:"employeeName"`
	Age      int
	State    string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling JSON data in Go")
	//EncodeJson()
	DecodeJSON()
}

func EncodeJson() {
	employees := []employee{
		{"Vinay Chunchu", 31, "Texas", "qwerty", []string{"java", "sql", "Go"}},
		{"V Chunchu", 31, "Texas", "ytrewq", []string{"python", "sql", "RUST"}},
		{"Vinay C", 31, "Texas", "qweytr", nil},
	}

	// Packaging above data to JSON data

	myJSON, err := json.MarshalIndent(employees, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", myJSON)

}

func DecodeJSON() {

	jsonDataResponse := []byte(`
		{
			"employeeName": "Vinay Chunchu",
			"Age": 31,
			"State": "Texas",
			"tags": ["java","sql","Go"]
		}
	`)

	var newEmployee employee

	isJsonValid := json.Valid(jsonDataResponse)

	if isJsonValid {
		fmt.Println("Provided JSON is valid")
		json.Unmarshal(jsonDataResponse, &newEmployee)
		fmt.Printf("%#v \n", newEmployee)
	} else {
		fmt.Println("Provided JSON is NOT valid")
	}

	// Example for the case where you want to add key value pairs
	// when dealing with http response which has key value pairs, key will be string, value could be interface{}

	var employeeMap map[string]interface{}

	json.Unmarshal(jsonDataResponse, &employeeMap)
	fmt.Printf("%#v \n", employeeMap)

	for key, value := range employeeMap {
		fmt.Printf("Key : %v , value : %v and Type : %T \n", key, value, value)
	}

}

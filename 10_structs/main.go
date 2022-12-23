package main

import "fmt"

func main() {
	fmt.Println("Structs learning")

	// There is no inheritence in GO lang. No super or parent as well.
	vinay := User{"VinayChunchu", "vinay@vinaychunchu.com", true, 31}
	fmt.Println(vinay)
	// %+v is used to get the property name of the struct printed out.
	fmt.Printf("Details of user vinay are : %+v \n", vinay)

	fmt.Printf("Name : %v \nEmail : %v \n", vinay.Name, vinay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

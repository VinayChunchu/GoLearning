package main

import "fmt"

func main() {
	fmt.Println("Pointers Learning")

	// * indicates that it is a pointer
	var intPointer *int
	fmt.Println("Initial value of pointer is : ", intPointer)

	someInt := 9

	// & is used for referencing the address of a variable
	var someIntPointer = &someInt
	fmt.Println("Address location of someInt : ", someIntPointer)
	fmt.Println("Value of someInt :", *someIntPointer)

	*someIntPointer = *someIntPointer * 9
	fmt.Println("Updated value of someInt: ", someInt)

}

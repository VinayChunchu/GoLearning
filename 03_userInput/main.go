package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Reading user input: "
	fmt.Println(welcome)

	//
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Go Languauge: ")

	// comma ok // err ok
	// Concept of try catch
	input, _ := reader.ReadString('\n')

	fmt.Printf("Thanks for providing ratings: %v  \n", input)
	fmt.Printf("Type of input variable is : %T \n", input)

}

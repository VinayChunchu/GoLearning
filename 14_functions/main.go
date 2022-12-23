package main

import "fmt"

func main() {
	fmt.Println("Functions Learning")
	callVinay()

	fmt.Println(addTwoNumbers(3, 5))

	total, message := addAllGivenNumbers(9, 43, 56, 2, 5, 5, 345, 324)
	fmt.Println(message, total)

}

func callVinay() {
	fmt.Println("Hola Migos from Vinay")
}

func addTwoNumbers(number1 int, number2 int) int {
	return number1 + number2
}

// number... is the syntax used in GO to get slices as input
func addAllGivenNumbers(numbers ...int) (int, string) {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total, "Sum of all numbers is :"
}

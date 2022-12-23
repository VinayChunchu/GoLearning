package main

import "fmt"

func main() {
	fmt.Println("Loops learning")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// Way one

	fmt.Println("First type of FOR loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Way two
	fmt.Println("Second type of FOR loop")
	for i := range days {
		fmt.Println(days[i])
	}

	// Way Three
	fmt.Println("Third type of FOR loop")
	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	// Way Four
	someValue := 1
	for someValue < 10 {

		if someValue == 2 {
			// goto is used to call Labels
			goto callVinay
		}

		if someValue == 5 {
			someValue++
			continue
		}
		fmt.Println("Value : ", someValue)
		someValue++
	}

callVinay:
	fmt.Println("Vinay is Calling")

}

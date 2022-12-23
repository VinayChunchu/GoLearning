package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case Learning")

	// rand method is used for getting random number
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")

	case 2:
		fmt.Println("Dice value is 2")

	case 3:
		fmt.Println("Dice value is 3")

		// if you want to run the next case , you can use fallthrough in all the case statements
		fallthrough
	case 4:
		fmt.Println("Dice value is 4")

	case 5:
		fmt.Println("Dice value is 5")

	case 6:
		fmt.Println("Dice value is 6")

	default:
		fmt.Println("Default Case")
	}

}

package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	fmt.Println("rand, crypto and maths")

	// var myFirstNumber int = 2
	// var mySecondNumber float64 = 4.5

	// This syntax ignores the decimals
	//fmt.Println("Sum of 2 integrers: ", myFirstNumber+int(mySecondNumber))

	// random number using math.rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandonNumber, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myRandonNumber)

}

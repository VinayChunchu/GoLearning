package main

import "fmt"

func main() {
	// when you add defer, that line will be executed as the last step in that function
	// if you have multiple defer statements, they are executed in LIFO order
	defer fmt.Println("Amigos")

	defer fmt.Println("One")

	defer fmt.Println("Two")

	fmt.Println("Hola")

	understandingDefer()

}

func understandingDefer() {
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}

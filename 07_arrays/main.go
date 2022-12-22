package main

import "fmt"

func main() {
	fmt.Println("Array's learning")

	var shapeArray [20]string

	shapeArray[0] = "BOX"
	shapeArray[1] = "ROUND"
	shapeArray[10] = "TRIANGLE"

	fmt.Println("Printing shapeArray values :", shapeArray)
	fmt.Println("Printing shapeArray length :", len(shapeArray))

	var carsArrays = [10]string{"Nissan", "Audi", "BMW"}
	fmt.Println("Printing Car Array: ", carsArrays)
	fmt.Println("Printing Car Array length: ", len(carsArrays))

}

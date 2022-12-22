package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Learning")

	var shapeArray = []string{}

	// In slices we use append to add items to it

	shapeArray = append(shapeArray, "BOX", "TRIANGLE", "ROUND", "PENTAGON")

	fmt.Println("Printing shapeArray values", shapeArray)
	fmt.Printf("Printing shapeArray type : %T \n", shapeArray)
	fmt.Println("Printing shapeArray length :", len(shapeArray))

	// : syntax is used to get sub set of slice
	// 1:3 , when we give from 1 to 3. The higher range value is not inclusive.
	// which means number 3 value is not included. for 1:3 -> it will give 1,2 values
	shapeArray = append(shapeArray[1:3])
	fmt.Println(shapeArray)

	shapeArray = append(shapeArray[:3])
	fmt.Println(shapeArray)

	// make is used to initialize slices
	carsPrices := make([]int, 4)

	carsPrices[0] = 234
	carsPrices[1] = 945
	carsPrices[2] = 465
	carsPrices[3] = 385

	// this will lead to index out of range
	//carsPrices[4] = 2948

	carsPrices = append(carsPrices, 8753, 583, 4334)
	fmt.Println(carsPrices)

	// Sorting
	fmt.Println("are carPrices sorted: ", sort.IntsAreSorted(carsPrices))
	sort.Ints(carsPrices)
	fmt.Println(carsPrices)

	// removing values from slices
	var cources = []string{"react", "java", "go", "python", "perl"}
	fmt.Println(cources)
	index := 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}

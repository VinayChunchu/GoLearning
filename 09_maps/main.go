package main

import "fmt"

func main() {
	fmt.Println("MAPS Learning")

	carsMap := make(map[string]string)

	carsMap["Germany"] = "BMW"
	carsMap["India"] = "TATA"
	carsMap["Japan"] = "HONDA"

	fmt.Println("List of all cars: ", carsMap)
	fmt.Println("Car from JAPAN : ", carsMap["Japan"])

	delete(carsMap, "Germany")
	fmt.Println("List of all cars after deletion: ", carsMap)

	// How loops works
	for key, value := range carsMap {
		fmt.Printf("For country %v we have %v car availble \n", key, value)
	}

}

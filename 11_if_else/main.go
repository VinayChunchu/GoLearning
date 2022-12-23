package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("If & Else learning")

	var loginCount string
	var result string

	reader := bufio.NewReader(os.Stdin)

	loginCount, _ = reader.ReadString('\n')

	if loginCount == "10" {
		result = "Regular user"
	} else if loginCount != "10" {
		result = "Else if user"
	}

	fmt.Println(result)

}

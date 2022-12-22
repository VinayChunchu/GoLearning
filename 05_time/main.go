package main

import (
	"fmt"
	"time"
)

func main() {

	// https://pkg.go.dev/time#pkg-overview

	fmt.Println("Welcome to Time Learning in GO")

	// present time
	// For formatting you have to use below date as input to change it !!!!!! not sure why :-)
	fmt.Println(time.Now().Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1991, time.March, 9, 11, 52, 30, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

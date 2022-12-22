package main

import "fmt"

// Constants
// when the first letter is in CAPS , it is global variable similar to "public" access modifier
const LoginToken string = "eddkkjhjbk839ljbn#DJKkh3kHB@3842JKHHBb"

func main() {
	var userName string = "Vinay"
	fmt.Println(userName)
	fmt.Printf("Type of userName variables : %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn variables : %T \n", isLoggedIn)

	/*
		https://go.dev/ref/spec#Numeric_types
		Numeric Types
		uint8       the set of all unsigned  8-bit integers (0 to 255)
		uint16      the set of all unsigned 16-bit integers (0 to 65535)
		uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
		uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		int8        the set of all signed  8-bit integers (-128 to 127)
		int16       the set of all signed 16-bit integers (-32768 to 32767)
		int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32     the set of all IEEE-754 32-bit floating-point numbers
		float64     the set of all IEEE-754 64-bit floating-point numbers

		complex64   the set of all complex numbers with float32 real and imaginary parts
		complex128  the set of all complex numbers with float64 real and imaginary parts

		byte        alias for uint8
		rune        alias for int32

	*/
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Type of smallInt variables : %T \n", smallInt)

	var smallFloat float32 = 255
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat variables : %T \n", smallFloat)

	// Alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of anotherVariable variables : %T \n", anotherVariable)

	//implicite Type
	var websiteUrl = "www.vinaychunchu.com"
	fmt.Println(websiteUrl)

	// No var keyword way of adding varialbles
	// This is only allowed inside methods
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Type of LoginTOken variables : %T \n", LoginToken)

}

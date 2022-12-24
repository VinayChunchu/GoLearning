package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const webURL = "https://google.com"

func main() {
	fmt.Println("Learning to make web requests")

	response, err := http.Get(webURL)

	checkNilError(err)

	fmt.Printf("Response is of type : %T \n", response)

	// *** It is caller's responsibility to close the connection
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(dataBytes)

	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

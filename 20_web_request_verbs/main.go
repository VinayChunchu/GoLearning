package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Using https://httpbin.org/ to perfrom diff http calls
// Thanks to Kenneth Reitz

func main() {
	fmt.Println("Learning various HTTP method verbs")

	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostJsonFormRequest()

}

func PerformGetRequest() {
	const myURL = "https://httpbin.org/get"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	// Another way to handle JSON's using string Builder

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count : ", byteCount)
	fmt.Println("Response data in string format using String Builder", responseString.String())

}

func PerformPostJsonRequest() {

	const myURL = "https://httpbin.org/post"

	// Some sample json data
	requestBody := strings.NewReader(`
		{
			"requester" : "Leaner of a GO Language",
			"requestType" : "POST",
			"requestLocation" : "Texas"
		}
	`)

	responseBody, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer responseBody.Body.Close()

	content, _ := ioutil.ReadAll(responseBody.Body)

	fmt.Println(string(content))

}

func PerformPostJsonFormRequest() {

	const myURL = "https://httpbin.org/post"

	//form data
	urlData := url.Values{}
	urlData.Add("requesterName", "Go learner")
	urlData.Add("requesterLocation", "Texas")
	urlData.Add("requesterTimeDay", "Saturday")

	responseData, err := http.PostForm(myURL, urlData)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(responseData.Body)

	fmt.Println(string(content))

}

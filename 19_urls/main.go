package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://api.open-meteo.com/v1/forecast?latitude=32.81&longitude=-96.95&hourly=temperature_2m"

func main() {
	fmt.Println("Learning how to handle url's")

	fmt.Println(myURL)

	// Parsing the URL

	parsingResult, _ := url.Parse(myURL)

	fmt.Println(parsingResult.Scheme)
	fmt.Println(parsingResult.Host)
	fmt.Println(parsingResult.Path)
	fmt.Println(parsingResult.Port())
	fmt.Println(parsingResult.RawQuery)

	queryParams := parsingResult.Query()

	fmt.Printf("Type of query Parameters is : %T \n", queryParams)

	fmt.Println(queryParams["latitude"])

	for _, values := range queryParams {
		fmt.Println(values)
	}

	// Constructing url if you have parts

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "api.open-meteo.com",
		Path:    "/v1/forecast",
		RawPath: "latitude=32.81&longitude=-96.95&hourly=temperature_2m",
	}

	constructedURL := partsOfURL.String()
	fmt.Println(constructedURL)

}

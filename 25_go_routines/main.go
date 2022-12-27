package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

// Wait groups are time.sleep() on steriods. These are used to manage the go routines
// Where ever we add them, it will wait for all the process to complete before closing main
var wg sync.WaitGroup // these are pointers

var mut sync.Mutex // pointer

func main() {

	// adding go keyword will create a seperate thread process
	// go greeter("Hell0")
	// greeter("World")

	websites := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		// this is tell run time how many go routines are there
		wg.Add(1)
	}

	// This is responsible for holdin main for all go routines to be completed
	wg.Wait()
	fmt.Println(signals)

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		// By waiting here, you are allowing the goroutine to complete before program exits
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {

	// This is responsible for saying when the go routine is done
	defer wg.Done()

	response, err := http.Get(endpoint)

	if err != nil {
		log.Println("Ran into Error !!!!!!!!")
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d status code for given url %s \n", response.StatusCode, endpoint)
}

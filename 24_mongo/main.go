package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VinayChunchu/GoLearning/mongo/router"
)

// Using mongo Atlas free shared cluster

func main() {
	fmt.Println("Integrating Mongo DB")

	fmt.Println("Server is starting....")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server Started. Listening on port 4000")
}

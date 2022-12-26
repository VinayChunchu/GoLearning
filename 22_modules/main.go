package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*

Go Modules Documentation: https://go.dev/ref/mod

Important Note : By default if you run go run main.go it will get dependencies from web or cache.

If you made any change to the vendor code by getting thier packages using "go mod vendor"

to use the updated code, you need to run you main function as "go run -mod=vendor main.go"

*/

func main() {
	fmt.Println("Learning more about modules")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello from greeter function")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is coming from serverHome function </h1>"))
}

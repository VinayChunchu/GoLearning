package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for our object: course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// Temporary DB
var courses []Course

// Temporary middleware or helpers - file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Learn how to build api's in GO")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "9", CourseName: "GO", CoursePrice: 10, Author: &Author{Fullname: "VinayChunchu", Website: "go.vinaychunchu.com"}})
	courses = append(courses, Course{CourseId: "10", CourseName: "JAVA", CoursePrice: 11, Author: &Author{Fullname: "VinayChunchu", Website: "java.vinaychunchu.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controller - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>JUST A SAMPLE HEADER RETURNING FROM GO LANG API CONTROLLER</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	// How to handle HTTP headers
	w.Header().Set("Content-Type", "application/json")
	// Handling JSON data in http calls
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get One courses")
	// How to handle HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Get id from request
	params := mux.Vars(r)

	// loop through our courses , find the matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// Return nothing found in case negitive scenario
	json.NewEncoder(w).Encode("No course found for the provided id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating One course")
	// How to handle HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Validating request body is not empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request Body is not valid, It came as empty")
	}

	// Validating request body came as {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Request Body is not valid, It came as {}")
		return
	}

	// Generate a new ID and append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating One course")
	// How to handle HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// Validating request body is not empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request Body is not valid, It came as empty")
	}

	// Validating request body came as {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Request Body is not valid, It came as {}")
		return
	}

	// extract the ID from request
	params := mux.Vars(r)

	// loop through courses, remove the course that matched with request ID and add it back with updated values
	// if course is not present , update a temporary boolean to return appropriate message

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO : Add more validations and cases of negitive scenarios

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting One course")
	// How to handle HTTP headers
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through courses, remove the course that matched with request ID
	// if course is not present , update a temporary boolean to return appropriate message
	var isCoursePresent bool

	for index, course := range courses {
		if course.CourseId == params["id"] {
			isCoursePresent = true
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Provided course ID is deleted")
			break
		} else {
			isCoursePresent = false
		}
	}

	if !isCoursePresent {
		json.NewEncoder(w).Encode("Provided course ID is not present")
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Student - Our struct for all articles
type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
	Age   string `json:"age"`
}

// Students - return n number articles
var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Students Page!")
	fmt.Println("Endpoint Hit: homePage for Students")
}

func fetchAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchAllStudents")
	json.NewEncoder(w).Encode(Students)
}

func fetchSingleStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchSingleStudent")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, student := range Students {
		if student.ID == key {
			fmt.Println(student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func addNewStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addNewStudent")
	// get the body of our POST request
	// unmarshal this into a new Student struct
	// append this to our Students array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student Student
	json.Unmarshal(reqBody, &student)
	// update our global student array to include
	// our new Student
	Students = append(Students, student)

	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, student := range Students {
		if student.ID == id {
			Students = append(Students[:index], Students[index+1:]...)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", fetchAllStudents)
	myRouter.HandleFunc("/addstudent", addNewStudent).Methods("POST")
	myRouter.HandleFunc("/student/{id}", fetchSingleStudent).Methods("POST")
	myRouter.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Students = []Student{
		Student{ID: "1", Name: "Raja", Grade: "4th Grade", Age: "10"},
		Student{ID: "2", Name: "Sivan", Grade: "6th Grade", Age: "12"},
		Student{ID: "3", Name: "Kumar", Grade: "7th Grade", Age: "13"},
	}
	handleRequest()
}

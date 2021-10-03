package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var base string = "http://localhost:8080/"

func main() {
	fmt.Println("Starting the application...")

	HttpGetCourse()
	HttpPostJSON()
	HttpGetCourseID("9810")
	HttpGetStudentsFromCourse("Algorithms and Data Structures")
}

func HttpGetCourse() {
	// Get response from http method
	response, _ := http.Get(base + "courses")

	// Read the data from  the response
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

func HttpPostJSON() {
	// Create jsonData
	jsonData := map[string]string{"message": "test"}

	// Marshal json
	jsonValue, _ := json.Marshal(jsonData)

	// Get response from http method
	response, _ := http.Post(base+"courses", "application/json", bytes.NewBuffer(jsonValue))

	// Read the data from  the response
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

func HttpGetCourseID(id string) {
	response, _ := http.Get(base + "courses/" + id)

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

func HttpGetStudentsFromCourse(courseName string) {
	response, _ := http.Get(base + "courses/" + courseName + "/students")

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

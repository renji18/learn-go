package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"dario.cat/mergo"
	"github.com/gorilla/mux"
)

// models
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author,omitempty"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website,omitempty"`
}

type ResponsePayload struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func sendJson(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := ResponsePayload{
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func main() {
	fmt.Println("API - LEARN GO")

	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{CourseId: "1", CourseName: "Go Lang", CoursePrice: 100, Author: &Author{Name: "Renji", Website: "renjiriverstone.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN Stack", CoursePrice: 100, Author: &Author{Name: "Renji", Website: "mern.dev"}})

	api := r.PathPrefix("/api").Subrouter()

	// routing
	api.HandleFunc("/", serveHome).Methods("GET")
	api.HandleFunc("/courses", getAllCourses).Methods("GET")
	api.HandleFunc("/course/single/{courseId}", getSingleCourse).Methods("GET")
	api.HandleFunc("/course/add", addCourse).Methods("POST")
	api.HandleFunc("/course/update/{courseId}", updateCourse).Methods("PUT")
	api.HandleFunc("/course/delete/{courseId}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	sendJson(w, 200, "Courses fetched successfully", courses)
}

// get single course
func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseId := params["courseId"]

	if courseId == "" {
		sendJson(w, 403, "courseId not provided", nil)
		return
	}

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == courseId {
			sendJson(w, 200, "Course fetched successfully", course)
			return
		}
	}

	sendJson(w, 200, "No course found with id: "+courseId, nil)
}

// add course
func addCourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		sendJson(w, 404, "Please send some data", nil)
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		sendJson(w, 404, "Invalid json provided", nil)
		return
	}

	if course.IsEmpty() {
		sendJson(w, 404, "Please provide a course name", nil)
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	sendJson(w, 404, "Course added successfully", map[string]string{"courseId": course.CourseId})
}

// update course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseId := params["courseId"]

	if courseId == "" {
		sendJson(w, 404, "No courseId provided", nil)
		return
	}

	if r.Body == nil {
		sendJson(w, 404, "No body provided", nil)
		return
	}

	for i, course := range courses {
		if course.CourseId == courseId {
			var courseBody Course
			err := json.NewDecoder(r.Body).Decode(&courseBody)
			if err != nil {
				sendJson(w, 404, "Error parsing json", nil)
				return
			}

			err = mergo.Merge(&course, courseBody, mergo.WithOverride)
			if err != nil {
				sendJson(w, 404, "Error updating course", nil)
				return
			}

			courses[i] = course

			sendJson(w, 201, "Error updating course", nil)
			return
		}
	}

	sendJson(w, 404, "Course not found", nil)
}

// delete course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseId := params["courseId"]

	if courseId == "" {
		sendJson(w, 403, "No courseId provided", nil)
		return
	}

	for i, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:i], courses[i+1:]...)

			sendJson(w, 403, "Course deleted successfully", nil)
			return
		}
	}

	sendJson(w, 404, "Course not found", nil)
}

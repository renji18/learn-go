/*
	Task 1: Route based Responses
	Build a server with routes

	/           → "Welcome"
	/hello      → "Hello"
	/bye        → "Bye"

	Constraint:
		•	use ServeMux (not global)
		•	no loops, no abstraction

	What this fixes:
		•	routing clarity
		•	handler separation
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/bye", byeHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the root route handler")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the hello route handler")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the bye route handler")
}

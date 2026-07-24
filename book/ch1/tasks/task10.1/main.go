/*
	Task 10 — First Real Abstraction

	Create a custom handler type:
		type AppHandler struct{}

	And implement:
		func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)

	Route everything through this single handler:
		/hello → Hello
		/bye   → Bye
		/      → Welcome

	Constraints:
		•	DO NOT use multiple HandleFunc
		•	Use one handler
		•	Inside it → route manually
*/

package main

import (
	"fmt"
	"net/http"
)

type AppHandler struct{}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	switch url {
	case "/hello":
		fmt.Fprintf(w, "Hello")

	case "/bye":
		fmt.Fprintf(w, "Bye")

	case "/":
		fmt.Fprintf(w, "Welcome")

	default:
		fmt.Fprintf(w, "Unregistered path: %s\n", r.URL.Path)
	}
}

func main() {
	http.ListenAndServe("localhost:8000", AppHandler{})
}

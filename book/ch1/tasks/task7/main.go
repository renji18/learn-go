/*
Task 7 — Calculator (but real, not toy)

	Route:
		/calc?a=10&b=20&op=add

	Supported ops:
		add → a + b
		sub → a - b
		mul → a * b
		div → a / b

	Constraints (this is where you’ll struggle):
		•	Missing param → 400 Bad Request
		•	Invalid number → 400
		•	division by zero → 400
		•	unknown op → 400
		•	error messages must include param name
		•	no duplicate parsing logic
		•	no float unless required
		•	use constants for ops

	Output format:
		result = 30

	You MUST:
		•	create helper functions:
		func parseIntParam(r *http.Request, key string) (int, error)
		func parseOp(r *http.Request) (string, error)
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	opAdd = "add"
	opSub = "sub"
	opMul = "mul"
	opDiv = "div"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/calc", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	operator, err := parseOp(r)
	if handleError(w, err) {
		return
	}

	a, err := parseIntParam(r, "a")
	if handleError(w, err) {
		return
	}

	b, err := parseIntParam(r, "b")
	if handleError(w, err) {
		return
	}

	result, err := calculator(a, b, operator)
	if handleError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "result = %d", result)
}

func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

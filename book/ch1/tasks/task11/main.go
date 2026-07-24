/*
	Task 11 — Logging Wrapper

	Log every request:
		GET /hello
		POST /calc

	You must create:
		func loggingMiddleware(next http.Handler) http.Handler

	Behavior
		•	log method + path
		•	then call next handler

	Constraints:
		•	You MUST use:
			next.ServeHTTP(w, r)
		•	This is the first time you’ll see:
			function wrapping behavior
*/

package main

import (
	"log"
	"net/http"
)

type AppHandler struct{}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func main() {

	log.Fatal(http.ListenAndServe("localhost:8000", loggingMiddleware(AppHandler{})))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return next
}

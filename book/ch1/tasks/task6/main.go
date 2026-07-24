/*
	Task 6 — Method Gatekeeping (but stricter)

	Route:
		/data

	Behavior:
		•	GET → return "GET OK"
		•	POST → return "POST OK"
		•	anything else → return 405 Method Not Allowed

	Constraints:
		•	MUST use:
			•	http.Error(...)
		•	MUST set proper status code:
		•	MUST NOT just print message — respect HTTP
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

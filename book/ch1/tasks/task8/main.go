/*
	Task 8 — Header Inspector (but intentional)

	Route:
		/headers

		Behavior:
			Return ONLY these headers:
				User-Agent
				Accept
				Host

	Constraints:
		•	If header missing → explicitly say:
			•	<header> = not present
		•	DO NOT loop blindly over all headers
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/headers", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

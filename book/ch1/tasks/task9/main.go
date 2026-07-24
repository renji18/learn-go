/*
	Task 9 — Response Discipline

	Route:
		/text

		Behavior:
			Return plain text:
				Hello, this is plain text response

	Constraints:
		• MUST explicitly set::
			•	Content-Type: text/plain
		•	MUST NOT mix binary/text (you already made that mistake earlier)
		•	MUST NOT rely on defaults
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/text", handler)
	mux.HandleFunc("/json", jsonHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, this is plain text response")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "hello",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

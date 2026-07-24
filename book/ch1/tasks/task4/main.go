/*
	Task 4 — Query-driven behavior (refinement)

	Route:
		/lissajous?cycle=10

	Behavior:
		•	default cycle = 5
		•	override if query param exists

	Constraint:
		•	use r.URL.Query().Get
		•	handle invalid input safely

	What this fixes:
		•	parsing discipline
		•	error handling
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/lissajous", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {

	c, err := cycle(r)

	if err != nil {
		http.Error(w, "Invalid cycle value", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/gif")
	lissajous(w, c)
}

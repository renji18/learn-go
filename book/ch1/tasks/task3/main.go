/*
	Task 3 — Counter endpoint (introduces: state)

	Route:
		/visit

	Each time user hits it:
		Visit count: 1
		Visit count: 2

	New concept:
		shared variable across requests

	What this teaches:
		•	package-level variables
		•	state in server
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int = 0

var mu sync.Mutex

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/visit", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock() // This has a slight overhead in large systems, but we don't have to worry about that for now.

	count++
	c := count
	// mu.Unlock() // Instead of writing this, we can defer the command

	// We did not put this inside the session, because in case Fprintf panics, the system never unlocks.
	fmt.Fprintf(w, "Visit count %d\n", c)
}

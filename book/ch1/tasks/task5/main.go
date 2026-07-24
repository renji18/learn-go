/*
	Task 5 — Multi-value query (introduces: slice understanding)

	Route:
		/?tag=go&tag=backend&tag=api

	Output:
		tag = go
		tag = backend
		tag = api

	New concept:
		map[string][]string in real use
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	tags, ok := r.URL.Query()["tag"]

	if !ok {
		fmt.Fprintf(w, "No tags")
		return
	}

	if len(tags) == 0 {
		fmt.Fprintf(w, "No query param found for 'tag'")
		return
	}

	for _, v := range tags {
		fmt.Fprintf(w, "tag = %s\n", v)
	}
}

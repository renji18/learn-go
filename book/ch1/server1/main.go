package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", handler) // each request calls handler
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler1)
	mux.HandleFunc("/test", handler2)
	mux.HandleFunc("/test/this", handler3)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

// handler echoes the path cmponent of the request URL r
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, from %q\t", r.URL.Path)
}

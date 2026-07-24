package main

import (
	"log"
	"net/http"
	"sync"
)

var count = make(map[string]int)
var mu sync.Mutex

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

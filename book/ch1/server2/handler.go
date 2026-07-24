package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	var path = r.URL.Path
	count[path]++
	mu.Unlock()
	fmt.Fprintf(w, "Count recorded for %s\n", path)
}

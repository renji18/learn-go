package main

import (
	"fmt"
	"net/http"
)

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from %q\t", r.URL.Path)
}

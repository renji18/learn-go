package main

import (
	"fmt"
	"net/http"
)

func counter(w http.ResponseWriter, r *http.Request) {
	if len(count) == 0 {
		fmt.Fprintf(w, "No count recorded yet")
	}

	for k, v := range count {
		fmt.Fprintf(w, "%s\t%d\n", k, v)
	}
}

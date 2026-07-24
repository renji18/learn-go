package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// method := r.Method

	// if method == "" || method == "GET" {
	// 	fmt.Fprintf(w, "GET OK")
	// 	return
	// }

	// if method == "POST" {
	// 	fmt.Fprintf(w, "POST OK")
	// 	return
	// }

	// // w.WriteHeader(405)
	// // fmt.Fprintf(w, "%s Method Not Allowed", method)

	// http.Error(w, method+" Method Not Allowed", 405)

	w.Header().Set("Content-Type", "text/plain")

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET OK")

	case http.MethodPost:
		fmt.Fprintf(w, "POST OK")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

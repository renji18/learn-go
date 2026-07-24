package main

import (
	// "fmt"
	// "log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cycle := 5

	// if err := r.ParseForm(); err == nil {
	// 	if v := r.Form.Get("cycle"); v != "" {
	// 		if n, err := strconv.Atoi(v); err == nil {
	// 			cycle = n
	// 		}
	// 	}
	// }

	if v := r.URL.Query().Get("cycle"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cycle = n
		}
	}

	w.Header().Set("Content-Type", "image/gif")
	lissajous(w, cycle)

	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.Proto, r.URL)

	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }

	// fmt.Fprintf(w, "Host= %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
}

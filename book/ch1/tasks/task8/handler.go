package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()

	if userAgent == "" {
		fmt.Fprintf(w, "User-Agent from r.UserAgent = not present")
	}
	fmt.Fprintf(w, "User-Agent from r.UserAgent: %s\n", userAgent)

	host := r.Host

	if host == "" {
		fmt.Fprintf(w, "Host from r.Host = not present")
	}
	fmt.Fprintf(w, "Host from r.Host: %s\n", host)

	headers := r.Header

	accept := headers.Get("Accept")
	userAgentHeader := headers.Get("User-Agent")
	// hostHeader := headers.Get("Host")

	if accept == "" {
		fmt.Fprintf(w, "Accept from header = not present")
	} else {
		fmt.Fprintf(w, "Accept from header: %s\n", accept)
	}

	if userAgentHeader == "" {
		fmt.Fprintf(w, "User-Agent from header = not present")
	} else {
		fmt.Fprintf(w, "User-Agent from header: %s\n", userAgentHeader)
	}

	// if hostHeader == "" {
	// 	fmt.Fprintf(w, "Host from header = not present")
	// } else {
	// 	fmt.Fprintf(w, "Host from header: %s\n", hostHeader)
	// }
}

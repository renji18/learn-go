package main

import (
	"fmt"
	"net/http"
	"strings"
)

func parseOp(r *http.Request) (string, error) {
	val := r.URL.Query().Get("op")

	if val == "" {
		return "", fmt.Errorf("missing operator param")
	}

	return strings.ToLower(val), nil
}

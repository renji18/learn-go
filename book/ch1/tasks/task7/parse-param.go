package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseIntParam(r *http.Request, key string) (int, error) {
	val := r.URL.Query().Get(key)

	if val == "" {
		return 0, fmt.Errorf("missing param for %s", key)
	}

	num, err := strconv.Atoi(val)

	if err != nil {
		return 0, fmt.Errorf("invalid number provided for %s", key)
	}

	return num, nil
}

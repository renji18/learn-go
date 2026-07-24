package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func cycle(r *http.Request) (int, error) {
	cycle := 5

	if v := r.URL.Query().Get("cycle"); v != "" {

		val, err := strconv.Atoi(v)

		if err != nil {
			return 0, err
		}

		if val < 1 {
			return 0, fmt.Errorf("Cycle must be >=1")
		}

		cycle = val
	}

	return cycle, nil
}

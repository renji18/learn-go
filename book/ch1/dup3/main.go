package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for line := range strings.SplitSeq(string(data), "\n") {
			counts[line]++
		}

		for line, n := range counts {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

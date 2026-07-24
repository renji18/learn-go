package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Enter whatever you want to enter")
		nullVal := false
		countLines(os.Stdin, counts, &nullVal)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			printFileName := false
			countLines(f, counts, &printFileName)

			if printFileName {
				fmt.Println(arg)
			}

			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int, printFileName *bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++

		for line, n := range counts {
			if n > 1 {
				*printFileName = true
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

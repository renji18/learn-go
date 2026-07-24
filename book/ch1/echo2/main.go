package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for idx, val := range os.Args[1:] {
		fmt.Println("At index ", idx, " we have ", val)
		s += sep + val
		sep = " "
	}

	fmt.Println(s)
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(*n, *sep, "THIS PRINTS THE VALUES")
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func init() {
	fmt.Println("Init no. 2")
}

func init() {
	fmt.Println("Init no. 1")
}

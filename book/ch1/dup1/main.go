package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter what you want to say: ")
	var text string

	for input.Scan() {
		text = input.Text()
		break

		// counts[input.Text()]++

		// for line, n := range counts {
		// 	if n > 1 {
		// 		fmt.Printf("%d\t%s\n", n, line)
		// 	}
		// }
	}

	fmt.Println("Program says:", text)

}

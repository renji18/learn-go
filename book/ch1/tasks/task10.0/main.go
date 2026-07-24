/*
	Task 10.0 — Type and method

	Create a custom calculator type
	Initialize it and handle calculation using the type
*/

package main

import "fmt"

type Calculator struct{}

func (c Calculator) Add(a int64, b int64) {
	fmt.Println(a + b)
}

func (c Calculator) Sub(a int64, b int64) {
	fmt.Println(a - b)
}

func main() {
	c := Calculator{}

	c.Add(5, 6)
	c.Sub(3, 100)
}

package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}

	fmt.Println("data: ", data)
	data = nonEmpty(data)
	fmt.Println("non empty data: ", nonEmpty(data))
	fmt.Println("data: ", data)
}

func nonEmpty(data []string) []string {
	i := 0
	for _, val := range data {
		if val != "" {
			data[i] = val
			i++
		}
	}

	return data[:i]
}

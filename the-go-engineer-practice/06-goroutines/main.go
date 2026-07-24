package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	ansC := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		ansC <- "Hello"
	}()

	select {
	case <-timer.C:
		fmt.Println("Time over")
	case gAns := <-ansC:
		fmt.Println("Input from goroutine: ", gAns)
	}
}

// func main() {
// 	ansC := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)

// 		ansC <- "Hello"
// 	}()

// 	msg := <-ansC

// 	fmt.Println(msg, "from goroutine")
// }

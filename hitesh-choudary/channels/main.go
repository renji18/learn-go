package main

import "fmt"

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int)

	go func(myCh chan<- int) {
		myCh <- 5
		close(myCh)
	}(myCh)

	val, isChannelOpen := <-myCh
	fmt.Print(val, isChannelOpen)
}

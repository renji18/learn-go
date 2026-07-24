package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}
	mu := &sync.RWMutex{}
	var score = []int{0}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		defer mu.Unlock()
		fmt.Println("R1")
		mu.Lock()
		score = append(score, 1)
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		defer mu.Unlock()
		fmt.Println("R2")
		mu.Lock()
		score = append(score, 2)
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		defer mu.Unlock()
		fmt.Println("R3")
		mu.Lock()
		score = append(score, 3)
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		defer mu.RUnlock()
		fmt.Println("R4")
		mu.RLock()
		fmt.Println(score)
	}(wg, mu)

	wg.Wait()
	fmt.Println(score)
}

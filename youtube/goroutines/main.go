package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	// go greeter("hello")
	// greeter("world")

	websiteListe := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://renjiriverstone.dev",
	}

	for _, web := range websiteListe {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()

	fmt.Println(signals)

	// We can us select {} instead of using a waitgroup as well, IFF the requirement is to keep the main thread blocked forever.
	// select {}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error in endopint %s: %v\n", endpoint, err)
		return
	}

	lock.Lock()
	signals = append(signals, endpoint)
	lock.Unlock()

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

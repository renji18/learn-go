/*
	Task 2: Echo query params (refinement)

	Input:
	/?name=Aadarsh&city=Ahmedabad

	Output:
	name = Aadarsh
	city = Ahmedabad

	Constraint:
	•	use r.URL.Query()
	•	do NOT use ParseForm

	What this fixes:
		•	proper mental model of query vs form
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if len(q) == 0 {
		fmt.Fprintf(w, "No query params")
		return
	}

	// for k, v := range q {
	// 	fmt.Fprintf(w, "%s = %s\n", k, v[0])
	// }


	// BETTER HANDLING OF V[0] USING VALUES RANGE, INSTEAD OF ASSUMING V EXISTS.
	for k, values := range q {
		for _, v := range values {
			fmt.Fprintf(w, "%s = %s\n", k, v)
		}
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renji18/mongo-api/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := mux.NewRouter()

	router.MovieRouter(r)

	fmt.Println("Listening to port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))

}

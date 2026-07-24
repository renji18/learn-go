package router

import (
	"github.com/gorilla/mux"
	controller "github.com/renji18/mongo-api/controller/movie"
)

func MovieRouter(mainRouter *mux.Router) *mux.Router {
	movieRouter := mainRouter.PathPrefix("/api/movie").Subrouter()

	movieRouter.HandleFunc("/", controller.CreateMovie).Methods("POST")
	movieRouter.HandleFunc("/all", controller.GetAllMovies).Methods("GET")
	movieRouter.HandleFunc("/delete", controller.DeleteAllMovie).Methods("DELETE")
	movieRouter.HandleFunc("/{movieId}", controller.WatchMovie).Methods("GET")
	movieRouter.HandleFunc("/{movieId}", controller.DeleteMovie).Methods("DELETE")

	return movieRouter
}

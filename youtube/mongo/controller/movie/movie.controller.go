package controller

import (
	"encoding/json"
	"net/http"

	"github.com/renji18/mongo-api/model"
	"github.com/renji18/mongo-api/util"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	allMovies := getAllMovies()
	err := util.SendJson(w, 200, "Movies fetched successfully", allMovies)
	util.HandleError(err)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	if r.Body == nil {
		err := util.SendJson(w, 403, "No data provided", nil)
		util.HandleError(err)
	}

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	util.HandleError(err)

	if movie.Movie == "" {
		err := util.SendJson(w, 403, "Movie name is required", nil)
		util.HandleError(err)
		return
	}

	id := insertOneMovie(movie)

	err = util.SendJson(w, 201, "Movie added successfully", map[string]any{"movieId": id})
	util.HandleError(err)
}

func WatchMovie(w http.ResponseWriter, r *http.Request) {
	movieId := util.GetParam(w, r, "movieId")

	watched := watchMovie(movieId)

	if !watched {
		err := util.SendJson(w, 404, "No records found for id "+movieId, nil)
		util.HandleError(err)
		return
	}

	err := util.SendJson(w, 201, "Movie watched successfully", nil)
	util.HandleError(err)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieId := util.GetParam(w, r, "movieId")

	deleted := deleteMovie(movieId)

	if !deleted {
		err := util.SendJson(w, 404, "No records found for id "+movieId, nil)
		util.HandleError(err)
		return
	}

	err := util.SendJson(w, 200, "Movie deleted successfully", nil)
	util.HandleError(err)
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	deleteAllMovie()
	err := util.SendJson(w, 200, "All movies deleted successfully", nil)
	util.HandleError(err)
}

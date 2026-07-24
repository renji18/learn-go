package controller

import (
	"context"

	"github.com/renji18/mongo-api/db"
	"github.com/renji18/mongo-api/model"
	"github.com/renji18/mongo-api/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// insert 1 movie
func insertOneMovie(movie model.Netflix) any {
	res, err := db.Collection.InsertOne(context.Background(), movie)
	util.HandleError(err)

	return res.InsertedID
}

// watch movie
func watchMovie(movieId string) bool {
	id := util.GetObjectId(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := db.Collection.UpdateOne(context.Background(), filter, update)
	util.HandleError(err)

	if res.MatchedCount == 0 {
		return false
	}
	return true
}

// delete 1 movie
func deleteMovie(movieId string) bool {
	id := util.GetObjectId(movieId)

	filter := bson.M{"_id": id}

	res, err := db.Collection.DeleteOne(context.Background(), filter)
	util.HandleError(err)

	if res.DeletedCount == 0 {
		return false
	}
	return true
}

// delete all movie
func deleteAllMovie() {
	_, err := db.Collection.DeleteMany(context.Background(), bson.D{{}})
	util.HandleError(err)
}

// get all movies
func getAllMovies() []bson.M {
	cursor, err := db.Collection.Find(context.Background(), bson.D{{}})
	util.HandleError(err)

	defer cursor.Close(context.Background())

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		util.HandleError(err)
		movies = append(movies, movie)
	}

	return movies
}

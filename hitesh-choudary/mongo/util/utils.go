package util

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetObjectId(movieId string) bson.ObjectID {
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func GetParam(w http.ResponseWriter, r *http.Request, findParam string) string {
	params := mux.Vars(r)

	found := params[findParam]

	if found == "" {
		SendJson(w, 501, findParam+"not provided", nil)
		log.Fatal(findParam + "not provided")
	}

	return found
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package db

import (
	"fmt"
	"log"

	"github.com/renji18/mongo-api/util"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Most Important
var Collection *mongo.Collection

// Connect with mongoDB
func init() {
	// parse config file
	envConfig, err := util.ParseConfig("/Users/renji/Desktop/projects/learn-go/youtube/mongo/.conf")
	if err != nil {
		log.Fatal(err)
	}

	// client option
	clientOptions := options.Client().ApplyURI(envConfig.MONGO_URI)

	// connent to mongodb
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successful")

	// collection reference
	Collection = client.Database(envConfig.DB_NAME).Collection(envConfig.COLLECTION_NAME)
	fmt.Println("Collection reference is ready")
}

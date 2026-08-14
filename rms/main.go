package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/renji18/rms/database"
	"github.com/renji18/rms/modules/auth"
	"github.com/renji18/rms/modules/menu"
	"github.com/renji18/rms/modules/restaurant"
	"github.com/renji18/rms/redis"
	"github.com/renji18/rms/utils"
)

var migration = flag.Bool("m", false, "Run or validate schema migrations")

func main() {
	// parse cli flags and env configs
	flag.Parse()
	utils.ParseConfig()

	// initialize redis
	redis.InitializeRedis()
	defer redis.RedisClient.Close()

	// initialize database
	database.InitializeDB()
	defer database.DB.Close()

	// run migrations or validations
	if *migration {
		database.Migration()
	} else {
		database.ValidateMigration(false)
	}

	// initialize an empty mux
	mux := http.NewServeMux()
	appMux := http.NewServeMux()

	// Register routers
	auth.AuthRouter(appMux)
	restaurant.RestaurantRouter(appMux)
	menu.MenuRouter(appMux)

	// Register app mux in default mux
	mux.Handle("/api/", http.StripPrefix("/api", appMux))

	// Run server
	fmt.Printf("Server is listening on port: %s...\n", utils.Config.PORT)
	log.Fatal(http.ListenAndServe(":"+utils.Config.PORT, mux))
}

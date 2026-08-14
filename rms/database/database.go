package database

import (
	"database/sql"
	"fmt"

	// the pq library. It just needs to be initialized, so we add an _ before an import, which is not being used, but needs to be initialized
	_ "github.com/lib/pq"
	"github.com/renji18/rms/utils"
)

// the database instance
var DB *sql.DB

func InitializeDB() {
	var err error
	// open database connection (does not guarantee reachability)
	DB, err = sql.Open(utils.Config.DRIVER, utils.Config.DATABASE_URL)
	utils.Fatal(fmt.Errorf("Error opening db connection: %v", err), err)

	// make sure database is reachable
	err = DB.Ping()
	utils.Fatal(fmt.Errorf("Error reacing db: %v", err), err)

	fmt.Println("Database connected!!")
}

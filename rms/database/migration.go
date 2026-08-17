package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/renji18/rms/database/queries"
	"github.com/renji18/rms/utils"
)

var createTableIfNotExists = `CREATE TABLE IF NOT EXISTS schema_version (version INT NOT NULL);`
var selectVersionFromTable = `SELECT version FROM schema_version LIMIT 1;`
var insertInitialIntoTable = `INSERT INTO schema_version (version) VALUES (0);`
var updateVersionInTable = `UPDATE schema_version SET version = $1`

// DEVELOPER RESPONSIBILITY
// The Validation and Migration function only work as expected, if this number is accurate. If there are 7 levels of migration, but this variable is at 5, then only the first 5 migrations will be validated/applied, and the application will startup successfully, but the database will remain out of sync.
var schemaVersion = 8

// This fn applies and syncs migration
func Migration() {
	currentVersion := ValidateMigration(true)

	if currentVersion == schemaVersion {
		fmt.Println("All migrations success.")
		os.Exit(0)
	}

	switch currentVersion + 1 {
	case 1:
		_, err := DB.Exec(queries.Query1_Create_Restaurant)
		utils.Fatal(fmt.Errorf("Error applying first migration: %v", err), err)
		fmt.Println("Migrated to version 1")

	case 2:
		_, err := DB.Exec(queries.Query2_Menu_And_Table_Connection_In_Restaurant)
		utils.Fatal(fmt.Errorf("Error applying second migration: %v", err), err)
		fmt.Println("Migrated to version 2")

	case 3:
		_, err := DB.Exec(queries.Query3_Auth_And_Owner_Tables)
		utils.Fatal(fmt.Errorf("Error applying third migration: %v", err), err)
		fmt.Println("Migrated to version 3")

	case 4:
		_, err := DB.Exec(queries.Query4_Is_Admin_And_Email)
		utils.Fatal(fmt.Errorf("Error applying forth migration: %v", err), err)
		fmt.Println("Migrated to version 4")

	case 5:
		_, err := DB.Exec(queries.Query5_Refresh_Token_And_Restaurant_Owner)
		utils.Fatal(fmt.Errorf("Error applying fifth migration: %v", err), err)
		fmt.Println("Migrated to version 5")

	case 6:
		_, err := DB.Exec(queries.Query6_Remove_Temp_Fields_From_Auth)
		utils.Fatal(fmt.Errorf("Error applying sixth migration: %v", err), err)
		fmt.Println("Migrated to version 6")

	case 7:
		_, err := DB.Exec(queries.Query7_Toggle_MenuItem_Status)
		utils.Fatal(fmt.Errorf("Error applying seventh migration: %v", err), err)
		fmt.Println("Migrated to version 7")

	case 8:
		_, err := DB.Exec(queries.Query8_Change_Default_For_Item_Status)
		utils.Fatal(fmt.Errorf("Error applying eighth migration: %v", err), err)
		fmt.Println("Migrated to version 8")
	}

	_, err := DB.Exec(updateVersionInTable, schemaVersion)
	utils.Fatal(fmt.Errorf("Error updating migration: %v", err), err)

	fmt.Println("All migrations success.")
	os.Exit(0)
}

// Function used to run validation. If runMigration is false, it is being used for valiation and breaking the program. If runMigration is true, then this fn later returns the currentVersion to the Migration fn which applies migrations
func ValidateMigration(runMigration bool) int {
	_, err := DB.Exec(createTableIfNotExists)
	utils.Fatal(fmt.Errorf("Error creating schema_version table: %v", err), err)

	var currentVersion int

	if err = DB.QueryRow(selectVersionFromTable).Scan(&currentVersion); err == sql.ErrNoRows {
		_, err = DB.Exec(insertInitialIntoTable)
		utils.Fatal(fmt.Errorf("Error inserting in schema_version table: %v", err), err)
		currentVersion = 0
	} else {
		utils.Fatal(fmt.Errorf("Error creating schema_version table: %v", err), err)
	}

	if !runMigration && currentVersion != schemaVersion {
		utils.Fatal(fmt.Errorf("Schema version mismatch. Run migration to sync schema."), errors.New(""))
	}

	return currentVersion
}

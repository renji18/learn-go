package table

import (
	"database/sql"

	"github.com/renji18/rms/database"
	schema "github.com/renji18/rms/types"
)

func createTable(data schema.RTable, restaurantId string) (tableId, statusCode int, message string, success bool) {
	// validate if data not empty
	if data.IsEmpty() {
		return 0, 422, "Please provide table number and seat count", false
	}

	// insert in db
	if err := database.DB.QueryRow(`
			INSERT INTO r_table (table_number, chair_count, restaurant_id)
			VALUES ($1, $2, $3)
			RETURNING id;
		`, data.TableNumber, data.ChairCount, restaurantId).Scan(&tableId); err != nil {
		return 0, 500, "Error inserting in db: " + err.Error(), false
	}

	// return tableId
	return tableId, 201, "Table created successfully", true
}

func updateTable(data schema.RTable, restaurantId string) (statusCode int, message string) {
	// validate if data is not empty
	if data.IsEmpty() {
		return 422, "Please provide table number and seat count"
	}

	// update in db
	if res, err := database.DB.Exec(`
			UPDATE r_table
			SET table_number = $3, chair_count = $4
			WHERE id = $1 AND restaurant_id = $2;
		`, data.ID, restaurantId, data.TableNumber, data.ChairCount); err != nil {
		return 500, "Error updating in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "Table not found"
	}

	return 200, "Table updated successfully"
}

func getAllTables(restaurantId string) (items []schema.RTable, statusCode int, message string, success bool) {
	// get all from database
	rows, err := database.DB.Query(`
			SELECT id, table_number, chair_count, occupied, reserved
			FROM r_table
			WHERE id = $1;
		`, restaurantId)
	if err != nil {
		var errMessage = "Error fetching items from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "Restaurant with id " + restaurantId + " not found"
		}

		return nil, 500, errMessage, false
	}

	defer rows.Close()

	for rows.Next() {
		var newTable schema.RTable
		err := rows.Scan(&newTable.ID, &newTable.TableNumber, &newTable.ChairCount, &newTable.Occupied, &newTable.Reserved)
		if err != nil {
			return nil, 500, "Error reading row: " + err.Error(), false
		}

		items = append(items, newTable)
	}

	return items, 200, "Tables fetched successfully", true
}

func getAllAvailableTables(restaurantId string) (items []schema.RTable, statusCode int, message string, success bool) {
	// get all from database
	rows, err := database.DB.Query(`
			SELECT id, table_number, chair_count
			FROM r_table
			WHERE id = $1 AND occupied = false AND reserved = false;
		`, restaurantId)
	if err != nil {
		var errMessage = "Error fetching items from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "Restaurant with id " + restaurantId + " not found"
		}

		return nil, 500, errMessage, false
	}

	defer rows.Close()

	for rows.Next() {
		var newTable schema.RTable
		err := rows.Scan(&newTable.ID, &newTable.TableNumber, &newTable.ChairCount)
		if err != nil {
			return nil, 500, "Error reading row: " + err.Error(), false
		}

		items = append(items, newTable)
	}

	return items, 200, "Tables fetched successfully", true
}

func deleteTable(restaurantId, tableId string) (statusCode int, message string) {
	// delete from db
	if res, err := database.DB.Exec(`
			DELETE FROM r_table
			WHERE id = $1 AND restaurant_id = $2;
		`, tableId, restaurantId); err != nil {
		return 500, "Error deleting record from db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "No such table found for this restaurant"
	}

	return 200, "Table deleted successfully"
}

func occupyTable() {}

func freeTable() {}

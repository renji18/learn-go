package menu

import (
	"database/sql"

	"github.com/renji18/rms/database"
	schema "github.com/renji18/rms/types"
)

// Currently we just take the restaurantId string, when we do implement authentication and authorization, the restaurantId can be access from the req.user itself.
func addMenuItem(body schema.MenuItem, restaurantId string) (id, statusCode int, message string, success bool) {
	// validate data
	if body.IsEmpty() {
		return 0, 422, "Please provide item name and category", false
	}

	// make db entry
	if err := database.DB.QueryRow(`
			INSERT INTO menu_item (name, category, price, restaurant_id)
			VALUES ($1, $2, $3, $4)
			RETURNING id;
		`, body.Name, body.Category, body.Price, restaurantId).Scan(&id); err != nil {
		return 0, 500, "Error inserting in db: " + err.Error(), false
	}

	// return id
	return id, 201, "Menu item created successfully", true
}

func updateMenuItem(body schema.MenuItem, restaurantId string) (statusCode int, message string) {
	// validate data
	if body.IsEmpty() {
		return 422, "Please provide item name and category"
	}

	// make db update
	if res, err := database.DB.Exec(`
			UPDATE menu_item
			SET name = $3, category = $4, price = $5
			WHERE id = $1 AND restaurant_id = $2;
		`, body.ID, restaurantId, body.Name, body.Category, body.Price); err != nil {
		return 500, "Error updating record in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "No such menu item found for this restaurant"
	}

	return 200, "Item updated successfully"
}

func getAllMenuItem(restaurantId string) (flatItems []schema.MenuItem, groupedItems map[string][]schema.MenuItem, statusCode int, message string, success bool) {
	groupedItems = make(map[string][]schema.MenuItem)

	// get all from database
	rows, err := database.DB.Query(`
			SELECT id, name, category, price, available, created_at, restaurant_id
			FROM menu_item
			WHERE restaurant_id = $1
			ORDER BY created_at;
		`, restaurantId)
	if err != nil {
		var errMessage = "Error fetching items from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "Restaurant with id " + restaurantId + " not found"
		}

		return nil, nil, 500, errMessage, false
	}
	defer rows.Close()

	for rows.Next() {
		var newItem schema.MenuItem
		err := rows.Scan(&newItem.ID, &newItem.Name, &newItem.Category, &newItem.Price, &newItem.Available, &newItem.CreatedAt, &newItem.RestaurantId)
		if err != nil {
			return nil, nil, 500, "Error reading row: " + err.Error(), false
		}

		flatItems = append(flatItems, newItem)
	}

	for _, item := range flatItems {
		groupedItems[item.Category] = append(groupedItems[item.Category], item)
	}

	return flatItems, groupedItems, 200, "Items fetched successfully", true
}

func deleteMenuItem(restaurantId, itemId string) (statusCode int, message string) {
	// delete from db
	if res, err := database.DB.Exec(`
			DELETE FROM menu_item
			WHERE id = $1 AND restaurant_id = $2;
		`, itemId, restaurantId); err != nil {
		return 500, "Error deleting record from db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "No such menu item found for this restaurant"
	}

	return 200, "Item deleted successfully"
}

func toggleMenuItem(restaurantId, itemId string) (statusCode int, message string) {
	// find item in db
	var currentAvailableStatus sql.NullBool
	if err := database.DB.QueryRow(`
			SELECT available
			FROM menu_item
			WHERE restaurant_id = $1 AND id = $2
		`, restaurantId, itemId).Scan(&currentAvailableStatus); err != nil {
		return 500, "Error fetching item from db: " + err.Error()
	}

	if !currentAvailableStatus.Valid {
		return 404, "Item with id not found"
	}

	// update item in db

	if res, err := database.DB.Exec(`
			UPDATE menu_item
			SET available = $3
			WHERE restaurant_id = $1 AND id = $2
		`, restaurantId, itemId, !currentAvailableStatus.Bool); err != nil {
		return 500, "Error updating record in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "No such menu item found for this restaurant"
	}

	if currentAvailableStatus.Bool {
		message = "Item is now not available"
	} else {
		message = "Item is now available"
	}

	return 200, message
}

package restaurant

import (
	"database/sql"

	"github.com/renji18/rms/database"
	schema "github.com/renji18/rms/types"
)

func addNewRestaurant(body schema.Restaurant) (id, statusCode int, message string, success bool) {
	// validate data
	if body.IsEmpty() {
		return 0, 422, "Please provide a name", false
	}

	// make db entry
	if err := database.DB.QueryRow(`
			INSERT INTO restaurant (name)
			VALUES ($1)
			RETURNING id;
		`, body.Name).Scan(&id); err != nil {
		return 0, 500, "Error inserting in db: " + err.Error(), false
	}

	//  return id
	return id, 201, "Restaurant created successfully", true
}

func assignOwner(body AssignOwnerDto) (statusCode int, message string) {
	// validate body
	if body.IsEmpty() {
		return 422, "Please provide owner_id and restaurant_id"
	}

	// check if user exists in database
	var userId int
	if err := database.DB.QueryRow(`
			SELECT id
			FROM users
			WHERE id = $1;
		`, body.OwnerId).Scan(&userId); err != nil {
		var errMessage = "Error fetching restaurant from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "User not found"
		}

		return 500, errMessage
	}

	// check if restaurant already has an owner and return early if it is the same ownerid
	var oldOwnerId sql.NullInt64
	if err := database.DB.QueryRow(`
			SELECT owner_id
			FROM restaurant
			WHERE id = $1
		`, body.RestaurantId).Scan(&oldOwnerId); err != nil {
		if err != sql.ErrNoRows {
			return 500, "Error fetching restaurant in db: " + err.Error()
		}
	}

	if oldOwnerId.Valid {
		if int(oldOwnerId.Int64) == body.OwnerId {
			return 200, "Owner already assigned"
		} else {
			// remove current restaurantId from old owner
			if res, err := database.DB.Exec(`
			UPDATE users
			SET restaurant_id = null
			WHERE id = $1
		`, oldOwnerId.Int64); err != nil {
				return 500, "Error removing user as restaurant owner in db: " + err.Error()
			} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
				return 404, "User not found"
			}
		}
	}

	// assign owner to the restaurant and to user table
	if res, err := database.DB.Exec(`
			UPDATE restaurant
			SET owner_id = $1
			WHERE id = $2
		`, body.OwnerId, body.RestaurantId); err != nil {
		return 500, "Error assigning owner to restaurant in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "Restaurant not found"
	}

	if res, err := database.DB.Exec(`
			UPDATE users
			SET restaurant_id = $1
			WHERE id = $2
		`, body.RestaurantId, body.OwnerId); err != nil {
		return 500, "Error updating user as a restaurant owner in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 404, "User not found"
	}

	// return
	return 201, "Onwer assigned successfully"
}

func getMyRestaurant(ownerId string) (restaurant schema.Restaurant, statusCode int, message string, success bool) {
	// find in db
	if err := database.DB.QueryRow(`
			SELECT r.id, r.name, r.created_at, COUNT(DISTINCT m.id) AS menu_item_count, COUNT(DISTINCT r_t.id) AS r_table_count
			FROM restaurant r
			LEFT JOIN menu_item m ON r.id = m.restaurant_id
			LEFT JOIN r_table r_t ON r.id = r_t.restaurant_id
			WHERE r.owner_id = $1
			GROUP BY
				r.id,
				r.name,
				r.created_at;
		`, ownerId).Scan(&restaurant.Id, &restaurant.Name, &restaurant.CreatedAt, &restaurant.Count.MenuItem, &restaurant.Count.RTable); err != nil {
		var errMessage = "Error fetching restaurant from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "No restaurant found for this user"
		}

		return restaurant, 500, errMessage, false
	}

	// return restaurant
	return restaurant, 200, "Restaurant fetched successfully", true
}

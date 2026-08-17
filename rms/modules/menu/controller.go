package menu

import (
	"encoding/json"
	"net/http"

	"github.com/renji18/rms/middleware"
	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func AddMenuItem(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	var itemBody schema.MenuItem
	err = json.NewDecoder(r.Body).Decode(&itemBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided "+err.Error(), nil)
		return
	}

	if id, statusCode, message, isSuccess := addMenuItem(itemBody, claims.RestaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]int{"item_id": id})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

func UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	var itemBody schema.MenuItem
	err = json.NewDecoder(r.Body).Decode(&itemBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil)
		return
	}

	statusCode, message := updateMenuItem(itemBody, claims.RestaurantId)

	utils.SendJson(w, statusCode, message, nil)
}

func GetAllMenuItem(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	if flatItems, groupedItems, statusCode, message, isSuccess := getAllMenuItem(claims.RestaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]any{"menu_items": flatItems, "grouped_by_category": groupedItems})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

func DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	itemId, found := utils.GetParam(w, r, "itemId")
	if !found {
		return
	}

	statusCode, message := deleteMenuItem(claims.RestaurantId, itemId)

	utils.SendJson(w, statusCode, message, nil)
}

func ToggleMenuItem(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	itemId, found := utils.GetParam(w, r, "itemId")
	if !found {
		return
	}

	statusCode, message := toggleMenuItem(claims.RestaurantId, itemId)

	utils.SendJson(w, statusCode, message, nil)
}

package menu

import (
	"encoding/json"
	"net/http"

	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func AddMenuItem(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	restaurantId, found := utils.GetParam(w, r, "restaurantId")

	if !found {
		return
	}

	var itemBody schema.MenuItem
	err := json.NewDecoder(r.Body).Decode(&itemBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil, nil)
		return
	}

	if id, statusCode, message, isSuccess := addMenuItem(itemBody, restaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]int{"item_id": id}, nil)
	} else {
		utils.SendJson(w, statusCode, message, nil, nil)
	}
}

func UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	restaurantId, found := utils.GetParam(w, r, "restaurantId")
	if !found {
		return
	}

	itemId, found := utils.GetParam(w, r, "itemId")
	if !found {
		return
	}

	var itemBody schema.MenuItem
	err := json.NewDecoder(r.Body).Decode(&itemBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil, nil)
		return
	}

	statusCode, message := updateMenuItem(itemBody, restaurantId, itemId)

	utils.SendJson(w, statusCode, message, nil, nil)
}

func GetAllMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantId, found := utils.GetParam(w, r, "restaurantId")

	if !found {
		return
	}

	if items, statusCode, message, isSuccess := getAllMenuItem(restaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string][]schema.MenuItem{"menu_items": items}, nil)
	} else {
		utils.SendJson(w, statusCode, message, nil, nil)
	}
}

func DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantId, found := utils.GetParam(w, r, "restaurantId")
	if !found {
		return
	}

	itemId, found := utils.GetParam(w, r, "itemId")
	if !found {
		return
	}

	statusCode, message := deleteMenuItem(restaurantId, itemId)

	utils.SendJson(w, statusCode, message, nil, nil)
}

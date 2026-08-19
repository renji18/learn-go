package table

import (
	"encoding/json"
	"net/http"

	"github.com/renji18/rms/middleware"
	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func CreateTable(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie "+err.Error(), nil)
		return
	}

	var tableBody schema.RTable
	err = json.NewDecoder(r.Body).Decode(&tableBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided "+err.Error(), nil)
		return
	}

	if id, statusCode, message, isSuccess := createTable(tableBody, claims.RestaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]int{"table_id": id})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie "+err.Error(), nil)
		return
	}

	var tableBody schema.RTable
	err = json.NewDecoder(r.Body).Decode(&tableBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided "+err.Error(), nil)
		return
	}

	statusCode, message := updateTable(tableBody, claims.RestaurantId)
	utils.SendJson(w, statusCode, message, nil)

}

func GetAllTables(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	if tables, statusCode, message, isSuccess := getAllTables(claims.RestaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]any{"tables": tables})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

func GetAllAvailableTables(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	if tables, statusCode, message, isSuccess := getAllAvailableTables(claims.RestaurantId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]any{"tables": tables})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

func DeleteTable(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	tableId, found := utils.GetParam(w, r, "tableId")
	if !found {
		return
	}

	statusCode, message := deleteTable(claims.RestaurantId, tableId)

	utils.SendJson(w, statusCode, message, nil)
}

func OccupyTable(w http.ResponseWriter, r *http.Request) {
	tableId, found := utils.GetParam(w, r, "tableId")
	if !found {
		utils.SendJson(w, 502, "Missing tableId", nil)
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 500, "Unauthorized: Not logged in", nil)
		return
	}

	statusCode, message := occupyTable(tableId, claims.RestaurantId)

	utils.SendJson(w, statusCode, message, nil)
}

func ReserveTable(w http.ResponseWriter, r *http.Request) {
	tableId, found := utils.GetParam(w, r, "tableId")
	if !found {
		utils.SendJson(w, 502, "Missing tableId", nil)
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 500, "Unauthorized: Not logged in", nil)
		return
	}

	statusCode, message := reserveTable(tableId, claims.RestaurantId)

	utils.SendJson(w, statusCode, message, nil)
}

func FreeTable(w http.ResponseWriter, r *http.Request) {
	tableId, found := utils.GetParam(w, r, "tableId")
	if !found {
		utils.SendJson(w, 502, "Missing tableId", nil)
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 500, "Unauthorized: Not logged in", nil)
		return
	}

	statusCode, message := freeTable(tableId, claims.RestaurantId)

	utils.SendJson(w, statusCode, message, nil)
}

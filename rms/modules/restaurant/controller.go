package restaurant

import (
	"encoding/json"
	"net/http"

	"github.com/renji18/rms/middleware"
	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func AddNewRestaurant(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		utils.SendJson(w, 403, "Body not provided", nil, nil)
		return
	}

	var restaurantBody schema.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurantBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil, nil)
		return
	}

	if id, statusCode, message, isSuccess := addNewRestaurant(restaurantBody); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]int{"restaurant_id": id}, nil)
	} else {
		utils.SendJson(w, statusCode, message, nil, nil)
	}
}

func AssignOwner(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		utils.SendJson(w, 404, "Body not provided", nil, nil)
		return
	}

	var body AssignOwnerDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil, nil)
		return
	}

	statusCode, message := assignOwner(body)

	utils.SendJson(w, statusCode, message, nil, nil)
}

func GetMyRestaurant(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil, nil)
		return
	}

	if restaurant, statusCode, message, isSuccess := getMyRestaurant(claims.UserId); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]schema.Restaurant{"restaurant": restaurant}, nil)
	} else {
		utils.SendJson(w, statusCode, message, nil, nil)
	}
}

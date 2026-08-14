package restaurant

import (
	"net/http"

	"github.com/renji18/rms/middleware"
)

func RestaurantRouter(appMux *http.ServeMux) {
	appMux.Handle("POST /restaurant", middleware.SuperAdminRoute(AddNewRestaurant))

	appMux.Handle("POST /restaurant/assign-owner", middleware.SuperAdminRoute(AssignOwner))

	appMux.Handle("GET /restaurant", middleware.ProtectedRoute(GetMyRestaurant))
}

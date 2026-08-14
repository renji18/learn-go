package menu

import (
	"net/http"

	"github.com/renji18/rms/middleware"
)

func MenuRouter(appMux *http.ServeMux) {
	appMux.Handle("POST /menu/{restaurantId}", middleware.ProtectedRoute(AddMenuItem))

	appMux.Handle("PUT /menu/{restaurantId}/{itemId}", middleware.ProtectedRoute(UpdateMenuItem))

	appMux.Handle("GET /menu/{restaurantId}", middleware.ProtectedRoute(GetAllMenuItem))

	appMux.Handle("DELETE /menu/{restaurantId}/{itemId}", middleware.ProtectedRoute(DeleteMenuItem))
}

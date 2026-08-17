package menu

import (
	"net/http"

	"github.com/renji18/rms/middleware"
)

func MenuRouter(appMux *http.ServeMux) {
	appMux.Handle("POST /menu", middleware.ProtectedRoute(AddMenuItem))

	appMux.Handle("PUT /menu", middleware.ProtectedRoute(UpdateMenuItem))

	appMux.Handle("GET /menu", middleware.ProtectedRoute(GetAllMenuItem))

	appMux.Handle("DELETE /menu/{itemId}", middleware.ProtectedRoute(DeleteMenuItem))

	appMux.Handle("PATCH /menu/{itemId}", middleware.ProtectedRoute(ToggleMenuItem))
}

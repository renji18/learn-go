package table

import (
	"net/http"

	"github.com/renji18/rms/middleware"
)

func TableRouter(appMux *http.ServeMux) {
	appMux.Handle("POST /table", middleware.ProtectedRoute(CreateTable))

	appMux.Handle("PUT /table", middleware.ProtectedRoute(UpdateTable))

	appMux.Handle("GET /table", middleware.ProtectedRoute(GetAllTables))

	appMux.Handle("GET /table/available", middleware.ProtectedRoute(GetAllAvailableTables))

	appMux.Handle("GET /occupy-table/{tableId}", middleware.ProtectedRoute(OccupyTable))

	appMux.Handle("GET /free-table/{tableId}", middleware.ProtectedRoute(FreeTable))
	
	appMux.Handle("DELETE /table/{tableId}", middleware.ProtectedRoute(DeleteTable))
}

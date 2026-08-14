package auth

import (
	"net/http"

	"github.com/renji18/rms/middleware"
)

func AuthRouter(appMux *http.ServeMux) {
	appMux.HandleFunc("POST /auth/login", Login)

	appMux.Handle("POST /auth/verify-otp", middleware.UnverfiedProtectedRoute(VerifyOtp))

	appMux.HandleFunc("GET /auth/refresh-token", RefreshToken)

	appMux.HandleFunc("GET /auth/logout", Logout)

	appMux.Handle("GET /auth/user", middleware.ProtectedRoute(GetUser))

	appMux.HandleFunc("PUT /auth/user", OnboardUser)

	appMux.Handle("POST /auth/user", middleware.SuperAdminRoute(AddNewUser))
}

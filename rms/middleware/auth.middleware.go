package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/renji18/rms/redis"
	"github.com/renji18/rms/utils"
)

type contextKey string

const userCtxKey contextKey = "user_claims"

func AuthMiddleware(next http.Handler, partialCheck bool, superAdminRoute bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(utils.Config.ACCESS_COOKIE_NAME)
		if err != nil {
			utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil, nil)
			return
		}

		claims, err := utils.ParseToken(cookie.Value)
		if err != nil {
			utils.SendJson(w, 401, err.Error(), nil, nil)
			return
		}

		if !claims.IsVerified && !partialCheck {
			utils.SendJson(w, 401, "Please verify otp before proceeding", nil, nil)
			return
		}

		if !partialCheck {
			// get access_token from redis
			if accessTokenFromRedis, err := redis.GetRedis("access_token" + claims.UserId); err != nil {
				utils.SendJson(w, 401, err.Error(), nil, nil)
				return
			} else {
				if cookie.Value != accessTokenFromRedis {
					utils.SendJson(w, 401, "Invalid token provided", nil, nil)
					return
				}
			}
		}

		if superAdminRoute && !claims.IsAdmin {
			utils.SendJson(w, 401, "Unauthorized route", nil, nil)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaims(ctx context.Context) (claims *utils.JwtClaims, err error) {
	claimsValue := ctx.Value(userCtxKey)
	if claimsValue == nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	var ok bool
	claims, ok = claimsValue.(*utils.JwtClaims)
	if !ok {
		return nil, fmt.Errorf("Internal Server Error")
	}

	return claims, nil
}

// Helper to easily wrap individual functions
func ProtectedRoute(hf http.HandlerFunc) http.Handler {
	return AuthMiddleware(hf, false, false)
}

func SuperAdminRoute(hf http.HandlerFunc) http.Handler {
	return AuthMiddleware(hf, false, true)
}

// Helper for unverified protected routes
func UnverfiedProtectedRoute(hf http.HandlerFunc) http.Handler {
	return AuthMiddleware(hf, true, false)
}

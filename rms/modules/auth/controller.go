package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/renji18/rms/middleware"
	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	var body LoginDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil)
		return
	}

	accessToken, expirationTime, statusCode, message := login(body)

	utils.SetCookie(w, utils.Config.ACCESS_COOKIE_NAME, accessToken, expirationTime)

	utils.SendJson(w, statusCode, message, nil)
}

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		return
	}

	var body struct {
		Otp string `json:"otp"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil)
		return
	}

	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	accessToken, refreshToken, accessExpiry, refreshExpiry, statusCode, message, isSuccess := verifyOtp(body.Otp, claims)

	if isSuccess {
		utils.SetCookie(w, utils.Config.ACCESS_COOKIE_NAME, accessToken, accessExpiry)

		utils.SetCookie(w, utils.Config.REFRESH_COOKIE_NAME, refreshToken, refreshExpiry)
	}

	utils.SendJson(w, statusCode, message, nil)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(utils.Config.REFRESH_COOKIE_NAME)
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	claims, err := utils.ParseToken(cookie.Value)
	if err != nil {
		utils.SendJson(w, 401, err.Error(), nil)
		return
	}

	accessToken, refreshToken, accessExpiry, refreshExpiry, statusCode, message, isSuccess := refreshToken(cookie.Value, claims)

	if isSuccess {
		utils.SetCookie(w, utils.Config.ACCESS_COOKIE_NAME, accessToken, accessExpiry)

		utils.SetCookie(w, utils.Config.REFRESH_COOKIE_NAME, refreshToken, refreshExpiry)
	}

	utils.SendJson(w, statusCode, message, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.SetCookie(w, utils.Config.ACCESS_COOKIE_NAME, "", time.Now())
	utils.SetCookie(w, utils.Config.REFRESH_COOKIE_NAME, "", time.Now())

	utils.SendJson(w, 200, "Logout successful", nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r.Context())
	if err != nil {
		utils.SendJson(w, 401, "Error accessing cookie: "+err.Error(), nil)
		return
	}

	user, statusCode, message, isSuccess := getUser(claims.UserId)
	if !isSuccess {
		utils.SendJson(w, statusCode, message, nil)
		return
	}

	utils.SendJson(w, statusCode, message, map[string]schema.User{"user": user})
}

func OnboardUser(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		utils.SendJson(w, 403, "Body not provided", nil)
		return
	}

	var userBody LoginDto
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil)
		return
	}

	statusCode, message := onboardUser(userBody)
	utils.SendJson(w, statusCode, message, nil)
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	if emptyBody := utils.ValidateBody(w, r.Body); emptyBody {
		utils.SendJson(w, 403, "Body not provided", nil)
		return
	}

	var userBody NewUserDto
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		utils.SendJson(w, 404, "Invalid json provided"+err.Error(), nil)
		return
	}

	if id, statusCode, message, isSuccess := addNewUser(userBody); isSuccess {
		utils.SendJson(w, statusCode, message, map[string]int{"user_id": id})
	} else {
		utils.SendJson(w, statusCode, message, nil)
	}
}

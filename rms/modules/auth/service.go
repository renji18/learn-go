package auth

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/renji18/rms/database"
	"github.com/renji18/rms/mail"
	"github.com/renji18/rms/redis"
	schema "github.com/renji18/rms/types"
	"github.com/renji18/rms/utils"
)

func login(body LoginDto) (accessToken string, expirationTime time.Time, statusCode int, message string) {
	// validate data
	if body.IsEmpty() {
		return "", time.Now(), 422, "Please provide email and password"
	}

	var user schema.User
	// get user with email
	if err := database.DB.QueryRow(`
			SELECT u.id, u.email, u.is_admin, a.password, a.is_active
			FROM users u
			LEFT JOIN auth a ON u.id = a.user_id
			WHERE email = $1;
		`, body.Email).Scan(&user.Id, &user.Email, &user.IsAdmin, &user.Auth.Password, &user.Auth.IsActive); err != nil {
		var errMessage = "Error fetching user from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "User with email " + body.Email + " not found"
		}

		return "", time.Now(), 500, errMessage
	}

	if !user.Auth.IsActive {
		return "", time.Now(), 401, "Please verify your account before proceeding ahead."
	}

	// authenticate password
	if err := utils.CheckPassword(body.Password, user.Auth.Password); err != nil {
		return "", time.Now(), 401, "Invalid credentials"
	}

	// generate access token (5 minutes)
	expirationTime = time.Now().Add(time.Minute * 5)

	claims := &utils.JwtClaims{
		Name:       user.Name,
		UserId:     user.Id,
		IsAdmin:    user.IsAdmin,
		IsVerified: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := utils.GetNewToken(claims)
	if err != nil {
		return "", time.Now(), 500, err.Error()
	}

	// generate otp
	otp := utils.GenerateOtp()
	fmt.Println(otp, "otp")

	// go store otp in redis
	err = redis.SetRedis("otp"+user.Id, otp, time.Minute*5)
	if err != nil {
		return "", time.Now(), 500, err.Error()
	}

	// go send email
	go mail.SendEmail([]string{"aadarsh@growception.com"}, "Your OTP for RMS",
		"<body>Your OTP for logging in to RMS is <b>"+otp+"</b><br />Please do not share this otp with anyone.<br /> Regards,<br />RMS Agent</body>")

	return accessToken, expirationTime, 201, "Otp sent successfully"
}

func verifyOtp(otp string, claims *utils.JwtClaims) (accessToken, refreshToken string, accessExpiry, refreshExpiry time.Time, statusCode int, message string, success bool) {
	// get otp from redis
	otpFromRedis, err := redis.GetRedis("otp" + claims.UserId)
	if err != nil {
		return "", "", time.Now(), time.Now(), 401, err.Error(), false
	}

	// verify otp (setup redis here to get otp)
	if otp != otpFromRedis {
		return "", "", time.Now(), time.Now(), 401, "Invalid OTP", false
	}

	err = redis.DelRedis("otp" + claims.UserId)
	if err != nil {
		return "", "", time.Now(), time.Now(), 401, err.Error(), false
	}

	// create and validate jwt tokens
	accessExpiry = time.Now().Add(time.Minute * 15)
	refreshExpiry = time.Now().Add(time.Hour * 24 * 30)

	acessClaim := &utils.JwtClaims{
		Name:       claims.Name,
		UserId:     claims.UserId,
		IsAdmin:    claims.IsAdmin,
		IsVerified: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaim := &utils.JwtClaims{
		UserId: claims.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err = utils.GetNewToken(acessClaim)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error generating access token: " + err.Error(), false
	}

	refreshToken, err = utils.GetNewToken(refreshClaim)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error generating refresh token: " + err.Error(), false
	}

	// store access token in redis
	err = redis.SetRedis("access_token"+claims.UserId, accessToken, time.Minute*30)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, err.Error(), false
	}

	// store refreshToken in db
	if res, err := database.DB.Exec(`
			UPDATE auth
			SET refresh_token = $2
			WHERE user_id = $1;
		`, claims.UserId, refreshToken); err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error storing refresh token in db: " + err.Error(), false
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return "", "", time.Now(), time.Now(), 404, "User not found", false
	}

	return accessToken, refreshToken, accessExpiry, refreshExpiry, 200, "Logged in successfully", true
}

func refreshToken(refreshTokenString string, claims *utils.JwtClaims) (accessToken, refreshToken string, accessExpiry, refreshExpiry time.Time, statusCode int, message string, success bool) {
	// get refresh token from db
	var refreshTokenFromDb sql.NullString

	if err := database.DB.QueryRow(`
			SELECT a.refresh_token
			FROM users u
			LEFT JOIN auth a ON u.id = a.user_id
			WHERE u.id = $1
		`, claims.UserId).Scan(&refreshTokenFromDb); err != nil {
		var errMessage = "Error fetching restaurant from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "User not found"
		}

		return "", "", time.Now(), time.Now(), 500, errMessage, false
	}

	if !refreshTokenFromDb.Valid {
		return "", "", time.Now(), time.Now(), 500, "No refresh token found.", false
	}

	// validate if it is correct
	if refreshTokenFromDb.String != refreshTokenString {
		return "", "", time.Now(), time.Now(), 500, "Invalid refresh token", false
	}

	// generate new access and new refresh token, similar to as in verify otp
	accessExpiry = time.Now().Add(time.Minute * 30)
	refreshExpiry = time.Now().Add(time.Hour * 24 * 30)

	acessClaim := &utils.JwtClaims{
		Name:       claims.Name,
		UserId:     claims.UserId,
		IsAdmin:    claims.IsAdmin,
		IsVerified: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaim := &utils.JwtClaims{
		UserId: claims.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	var err error
	accessToken, err = utils.GetNewToken(acessClaim)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error generating access token: " + err.Error(), false
	}

	refreshToken, err = utils.GetNewToken(refreshClaim)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error generating refresh token: " + err.Error(), false
	}

	// store new access token in redis
	err = redis.SetRedis("access_token"+claims.UserId, accessToken, time.Minute*30)
	if err != nil {
		return "", "", time.Now(), time.Now(), 500, err.Error(), false
	}

	// store new refresh token in db
	if res, err := database.DB.Exec(`
			UPDATE auth
			SET refresh_token = $2
			WHERE user_id = $1;
		`, claims.UserId, refreshToken); err != nil {
		return "", "", time.Now(), time.Now(), 500, "Error storing refresh token in db: " + err.Error(), false
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return "", "", time.Now(), time.Now(), 404, "User not found", false
	}

	// return new access and new refresh token
	return accessToken, refreshToken, accessExpiry, refreshExpiry, 201, "Session refreshed successfully", true
}

func getUser(userId string) (user schema.User, statusCode int, message string, success bool) {
	// find in db
	var restaurantName sql.NullString

	if err := database.DB.QueryRow(`
			SELECT u.id, u.name, u.email, u.is_admin, r.name
			FROM users u
			LEFT JOIN auth a ON u.id = a.user_id
			LEFT JOIN restaurant r ON u.id = r.owner_id
			WHERE u.id = $1 AND a.is_active = true;
		`, userId).Scan(&user.Id, &user.Name, &user.Email, &user.IsAdmin, &restaurantName); err != nil {
		var errMessage = "Error fetching user from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "User with id " + userId + " not found"
		}

		return user, 500, errMessage, false
	}

	if restaurantName.Valid {
		user.Restaurant.Name = restaurantName.String
	}

	// return user
	return user, 200, "User fetched successfully", true
}

func onboardUser(body LoginDto) (statusCode int, message string) {
	// validate body
	if body.IsEmpty() {
		return 404, "Please provide email and temp password shared in mail."
	}

	var userId string
	// get user from db
	if err := database.DB.QueryRow(`
			SELECT u.id
			FROM users u
			WHERE u.email = $1
		`, body.Email).Scan(&userId); err != nil {
		var errMessage = "Error fetching restaurant from db: " + err.Error()

		if err == sql.ErrNoRows {
			errMessage = "User not found"
		}

		return 500, errMessage
	}

	// get temp pw from redis
	if tempPwFromRedis, err := redis.GetRedis("temp_pw" + userId); err != nil {
		return 401, err.Error()
	} else {
		if body.Password != tempPwFromRedis {
			return 401, "Invalid password provided."
		}
	}

	// update user auth in db
	if res, err := database.DB.Exec(`
			UPDATE auth
			SET is_active = true
			WHERE user_id = $1
		`, userId); err != nil {
		return 500, "Error updating user auth in db: " + err.Error()
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 500, "User with provided email does not exist"
	}

	// return
	return 200, "User onboarded successfully"
}

func addNewUser(body NewUserDto) (id, statusCode int, message string, success bool) {
	// validate data
	if body.IsEmpty() {
		return 0, 422, "Please provide user name, email and pw", false
	}

	// check if user with email not exists
	var existingUserId sql.NullString
	if err := database.DB.QueryRow(`
			SELECT u.id
			FROM users u
			WHERE u.email = $1;
		`, body.Email).Scan(&existingUserId); err != nil {
		if err != sql.ErrNoRows {
			return 0, 500, "Error fetching user from db: " + err.Error(), false
		}
	}

	if existingUserId.Valid {
		return 0, 409, "User with this email already exists", false
	}

	// make db entry for user record
	if err := database.DB.QueryRow(`
			INSERT INTO users (name, email)
			VALUES ($1, $2)
			RETURNING id;
		`, body.Name, body.Email).Scan(&id); err != nil {
		return 0, 500, "Error inserting in db: " + err.Error(), false
	}

	// hash password
	hashedPw, err := utils.HashPassword(body.Password)
	if err != nil {
		return 0, 500, "Error hashing pw: " + err.Error(), false
	}

	// make db entry for auth
	if res, err := database.DB.Exec(`
			INSERT INTO auth (password, user_id)
			VALUES ($1, $2);
		`, hashedPw, id); err != nil {
		return 0, 500, "Error inserting in db: " + err.Error(), false
	} else if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return 0, 500, "Error creating auth record for user", false
	}

	// generate temp pw
	tempPw := utils.GenerateOtp()
	fmt.Println(tempPw, "tempPw")

	// store in redis
	err = redis.SetRedis("temp_pw"+strconv.Itoa(id), tempPw, 0)
	if err != nil {
		return 0, 500, err.Error(), false
	}

	// send temp pw in email
	go mail.SendEmail([]string{body.Email}, "Your Temporary Password for RMS", "<body>Your Temporary password for onboarding in to RMS is <b>"+tempPw+"</b><br/>This otp will expire in 1 hour.<br/>Please do not share this password with anyone.<br/>Regards,<br />RMS Agent</body>")

	// return id
	return id, 201, "User created successfully. Please onboard via the link in email.", true
}

package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Name       string `json:"name"`
	UserId     string `json:"user_id"`
	IsAdmin    bool   `json:"is_admin"`
	IsVerified bool   `json:"is_verified"`
	jwt.RegisteredClaims
}

func GetNewToken(claims *JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Config.JWT_SECRET))

	if err != nil {
		return "", fmt.Errorf("Error generating token: %v", err)
	}

	return tokenString, nil
}

func ParseToken(accessToken string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (any, error) {
		return []byte(Config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error verifying signature: %v", err)
	}

	if !tkn.Valid {
		return nil, fmt.Errorf("Invalid or expired token")
	}

	return claims, nil
}

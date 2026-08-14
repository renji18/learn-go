package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Error hashing password: %v\n", err)
	}

	return string(hashed), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateOtp() string {
	digits := "0123456789"
	var otp strings.Builder

	for range [6]int{} {
		otp.WriteByte(digits[rand.Intn(10)])
	}

	return otp.String()
}

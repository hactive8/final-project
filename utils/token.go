package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id uint, email string) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	claim := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return claims.SignedString([]byte(secret))
}

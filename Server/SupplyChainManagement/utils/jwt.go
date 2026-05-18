package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, email, role string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"iss":     "supply-chain-management",
		"iat":     jwt.NewNumericDate(time.Now()),
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}

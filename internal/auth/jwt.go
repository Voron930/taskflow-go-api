package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

func GenerateJWT(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (int64, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	userIDfloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	email, ok := claims["email"].(string)
	if !ok {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	return int64(userIDfloat), email, nil
}

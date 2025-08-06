package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/riteshco/Feasto/pkg/types"
)


func GenerateJWTToken(user types.User) (string , error) {

	secret := os.Getenv("JWT_SECRET")

	// Create claims
	claims := jwt.MapClaims{
		"id"      :  user.Id,
		"username":  user.Username,
		"email":     user.Email,
		"user_role": user.UserRole,
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),                     // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	
	return token.SignedString([]byte(secret))
}
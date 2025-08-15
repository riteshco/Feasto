package utils

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"os"
	"strings"
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

func IsValidEmail(email string) bool {
	parsed, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	addr := parsed.Address

	at := strings.LastIndex(addr, "@")
	if at < 0 {
		return false
	}
	domain := addr[at+1:]

	return strings.Contains(domain, ".") 
}

func ErrorHandling(w http.ResponseWriter ,message string , status int){
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(types.Message{Message: message})
}
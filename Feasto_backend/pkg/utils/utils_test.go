package utils

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/riteshco/Feasto/pkg/constants"
	"github.com/riteshco/Feasto/pkg/types"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"another@test.co", true},
		{"missing@domain", false},
	}

	for _, tt := range tests {
		result := IsValidEmail(tt.email)
		if result != tt.expected {
			t.Errorf("IsValidEmail(%q) = %v; want %v", tt.email, result, tt.expected)
		}
	}
}

func TestGenerateJWTToken(t *testing.T) {
	// Set environment variable for test
	secret := "testsecret"
	os.Setenv("JWT_SECRET", secret)

	user := types.User{
		Id:       123,
		Username: "testuser",
		Email:    "test@example.com",
		UserRole: constants.RoleAdmin,
	}

	tokenString, err := GenerateJWTToken(user)
	if err != nil {
		t.Fatalf("GenerateJWTToken() returned error: %v", err)
	}

	// Parse the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Fatalf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("Token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Failed to cast claims to MapClaims")
	}

	// Check some fields
	if claims["username"] != user.Username {
		t.Errorf("Expected username %q, got %q", user.Username, claims["username"])
	}
	if claims["email"] != user.Email {
		t.Errorf("Expected email %q, got %q", user.Email, claims["email"])
	}

	// Check expiry
	exp := int64(claims["exp"].(float64))
	if time.Unix(exp, 0).Before(time.Now()) {
		t.Error("Token expiry is in the past")
	}

	// Quick sanity check: token string should have 3 parts
	if len(strings.Split(tokenString, ".")) != 3 {
		t.Errorf("Token string %q is not a valid JWT format", tokenString)
	}
}
package passwords

import (
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "supersecret123"

	// Hash the password
	hashed := HashPassword(password)

	if hashed == "" {
		t.Fatal("expected hashed password to be non-empty")
	}

	if hashed == password {
		t.Fatal("hashed password should not match raw password")
	}

	if !VerifyHashPassword(password, hashed) {
		t.Error("expected password to match hash, but it did not")
	}

	if VerifyHashPassword("wrongpassword", hashed) {
		t.Error("expected wrong password to fail verification, but it succeeded")
	}
}

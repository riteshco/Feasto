package passwords

import "golang.org/x/crypto/bcrypt"



func HashPassword(password string) string {
	hashedPassword , _ := bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
	return string(hashedPassword)
}

func VerifyHashPassword(password string , hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
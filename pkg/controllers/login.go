package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
)

func AuthenticateUserAPI(w http.ResponseWriter , r *http.Request){

	var user types.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	username := user.Username
	email := user.Email
	password := user.Password

	if username == "" || email=="" || password==""{
		log.Fatalf("All fields are required to login!")
		toSend := types.Message{Message: "Empty fields"}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	dbUser , err := models.GetUserByEmail(r.Context() , email)
	if err != nil {
		log.Fatalf("Error fetching user from DB for authentication : %v" , err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "authentication successful",
		"user_id":  dbUser.Id,
		"username": dbUser.Username,
		"email":    dbUser.Email,
		"role":     dbUser.UserRole,
	})

	fmt.Println(dbUser.Id)
	fmt.Println(dbUser.Username)
	fmt.Println(dbUser.MobileNumber)
	fmt.Println(dbUser.Email)
	fmt.Println(dbUser.UserRole)
	fmt.Println(dbUser.HashedPassword)

}
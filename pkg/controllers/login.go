package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/riteshco/Feasto/pkg/controllers/passwords"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
	"github.com/riteshco/Feasto/pkg/utils"
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

	if ! passwords.VerifyHashPassword(password , dbUser.HashedPassword) {
		log.Fatalf("Wrong Password")
	}
	
	token , err := utils.GenerateJWTToken(dbUser)
	if err != nil {
		log.Fatalf("Error generating JWT Token : %v" , err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "authentication successful",
		"user_id":  dbUser.Id,
		"username": dbUser.Username,
		"email":    dbUser.Email,
		"role":     dbUser.UserRole,
		"token":    token,
	})

}

func AuthenticateUser(w http.ResponseWriter , r *http.Request){
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

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

	if ! passwords.VerifyHashPassword(password , dbUser.HashedPassword) {
		log.Fatalf("Wrong Password")
	}
	
	token , err := utils.GenerateJWTToken(dbUser)
	if err != nil {
		log.Fatalf("Error generating JWT Token : %v" , err)
	}

	// Save token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // set true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   60 * 60,
	})
	
}
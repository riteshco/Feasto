package controllers

import (
	"encoding/json"
	"fmt"
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
		toSend := types.Message{Message: `All three :- "username" , "email" and "password" fields are required to be non-empty to login.`}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}
	
	dbUser , status , err := models.GetUserByEmailDB(r.Context() , email)
	if err != nil {
		fmt.Printf("Error fetching user from DB for authentication : %v\n" , err)
		http.Error(w , err.Error() , status)
		return
	} else if ! passwords.VerifyHashPassword(password , dbUser.HashedPassword) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Wrong Password",
	})
	} else {
	
	token , err := utils.GenerateJWTToken(dbUser)
	if err != nil {
		fmt.Printf("Error generating JWT Token : %v" , err)
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

}

func AuthenticateUser(w http.ResponseWriter , r *http.Request){
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if username == "" || email=="" || password==""{
		toSend := types.Message{Message: `All three :- "username" , "email" and "password" fields are required to be non-empty to login.`}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	
	dbUser , status , err := models.GetUserByEmailDB(r.Context() , email)
	if err != nil {
		fmt.Printf("Error fetching user from DB for authentication : %v\n" , err)
		http.Error(w , err.Error() , status)
		return
	} else if ! passwords.VerifyHashPassword(password , dbUser.HashedPassword) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Wrong Password",
	})
	} else {
	
	token , err := utils.GenerateJWTToken(dbUser)
	if err != nil {
		fmt.Printf("Error generating JWT Token : %v" , err)
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
		
}
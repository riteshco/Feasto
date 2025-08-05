package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/riteshco/Feasto/pkg/controllers/passwords"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
)


func RegisterUser(w http.ResponseWriter , r *http.Request){
	username := r.PostFormValue("username")
	mobile_number := r.PostFormValue("mobile")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if username == "" || mobile_number== "" || email=="" || password==""{
		log.Fatalf("All fields are required to register!")
		toSend := types.Message{Message: "Empty fields"}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	hashed_password := passwords.HashPassword(password)

	register := types.UserRegisterDB{
		Username: username,
		MobileNumber: mobile_number,
		Email: email,
		UserRole: "customer",
		HashedPassword: hashed_password,
	}

	success , err := models.RegisterUser(register)
	if err != nil {
		log.Fatalf("Could not log user")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	if success {
		fmt.Println("User registered successfully")
		toSend := types.Message{Message: "User registered successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}
}

func RegisterAPIUser(w http.ResponseWriter , r *http.Request){
	var user types.UserRegister

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	username := user.Username
	mobile_number := user.MobileNumber
	email := user.Email
	password := user.Password

	if username == "" || mobile_number== "" || email=="" || password==""{
		log.Fatalf("All fields are required to register!")
		toSend := types.Message{Message: "Empty fields"}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	hashed_password := passwords.HashPassword(password)

	register := types.UserRegisterDB{
		Username: username,
		MobileNumber: mobile_number,
		Email: email,
		UserRole: "customer",
		HashedPassword: hashed_password,
	}

	success , err := models.RegisterUser(register)
	if err != nil {
		log.Fatalf("Could not log user")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	if success {
		fmt.Println("User registered successfully")
		toSend := types.Message{Message: "User registered successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			log.Fatal(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}

}
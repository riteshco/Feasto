package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/riteshco/Feasto/pkg/controllers/passwords"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
	"github.com/riteshco/Feasto/pkg/utils"
)


func RegisterUser(w http.ResponseWriter , r *http.Request){
	username := r.PostFormValue("username")
	mobile_number := r.PostFormValue("mobile")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	username = strings.TrimSpace(username)
	check_username := strings.ToLower(username)
	if check_username == "admin" || check_username == "chef" {
		fmt.Println("Tried to put admin or chef as username!")
		http.Error(w , "This username is not allowed!" , http.StatusBadRequest)
		return
	}

	if ! utils.IsValidEmail(email) {
		fmt.Println("Didn't enter a valid email!")
		http.Error(w , "Please enter a valid email!" , http.StatusBadRequest)
		return
	}

	if username == "" || mobile_number == "" || email == "" || password == "" {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(types.Message{
		Message: "All fields (username, mobile number, email, password) are required to register",
	})
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

	success , status , err := models.RegisterUserDB(register)
	if err != nil {
		fmt.Printf("Could not log user")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), status)
		return
	}
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered Successfully!!"))
		return
	}
}

func RegisterUserAPI(w http.ResponseWriter , r *http.Request){
	var user types.UserRegister

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	username := user.Username
	mobile_number := user.MobileNumber
	email := user.Email
	password := user.Password

	username = strings.TrimSpace(username)
	check_username := strings.ToLower(username)
	if check_username == "admin" || check_username == "chef" {
		fmt.Println("Tried to put admin or chef as username!")
		http.Error(w , "This username is not allowed!" , http.StatusBadRequest)
		return
	}

	if ! utils.IsValidEmail(email) {
		fmt.Println("Didn't enter a valid email!")
		http.Error(w , "Please enter a valid email!" , http.StatusBadRequest)
		return
	}

	if username == "" || mobile_number == "" || email == "" || password == "" {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(types.Message{
		Message: "All fields (username, mobile number, email, password) are required to register",
	})
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

	success , status , err := models.RegisterUserDB(register)
	if err != nil {
		fmt.Printf("Could not log user")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), status)
		return
	}
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User registered Successfully!!"))
		return
	}

}
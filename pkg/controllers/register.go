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

func validateSignupRequest (user  types.UserRegister ) error {
	if user.Username == "" || user.MobileNumber == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("all fields (username, mobile number, email, password) are required to register")
	}
	if len(user.MobileNumber) != 10 {
		return fmt.Errorf("length of mobile_number should be equal to 10")
	}
	if len(user.Username) > 12 {
		return fmt.Errorf("username length should be less than equal to 12")
	}
	if len(user.Password) < 6 || len(user.Password) > 20 {
		return fmt.Errorf("length of password should be between 6 to 20")
	}
	user.Username = strings.TrimSpace(user.Username)
	check_username := strings.ToLower(user.Username)
	if check_username == "admin" || check_username == "chef" {
		fmt.Println("Tried to put admin or chef as username!")
		return fmt.Errorf("this username is not allowed")
	}
	if ! utils.IsValidEmail(user.Email) {
		fmt.Println("Didn't enter a valid email!")
		return fmt.Errorf("please enter a valid email")
	}

	return nil
}


func RegisterUser(w http.ResponseWriter , r *http.Request){
	username := r.PostFormValue("username")
	mobile_number := r.PostFormValue("mobile")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := types.UserRegister{
		Username: username,
		MobileNumber: mobile_number,
		Email: email,
		Password: password,
	}

	err := validateSignupRequest(user)
	if err != nil {
		http.Error(w , err.Error() , http.StatusBadRequest)
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

	err := validateSignupRequest(user)
	if err != nil {
		http.Error(w , err.Error() , http.StatusBadRequest)
		return
	} 

	username := user.Username
	mobile_number := user.MobileNumber
	email := user.Email
	password := user.Password

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
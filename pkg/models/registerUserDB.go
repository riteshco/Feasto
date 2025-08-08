package models

import (
	"fmt"
	"net/http"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/riteshco/Feasto/pkg/types"
)

func RegisterUser(user types.UserRegisterDB) (bool , int , error) {


	InsertUser := "INSERT INTO Users (username , mobile_number , email , user_role , password_hash) VALUES (? , ? , ? , ? , ?)"
	_ , err := DB.Exec(InsertUser , user.Username , user.MobileNumber , user.Email , user.UserRole , user.HashedPassword)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate entry in registration")
			return false, http.StatusAlreadyReported , fmt.Errorf("user already exists")
		} else {
			fmt.Println("error inserting into the database", err)
			return false, http.StatusInternalServerError , fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("User registered successfully")
		return true, http.StatusOK , nil
	}
}
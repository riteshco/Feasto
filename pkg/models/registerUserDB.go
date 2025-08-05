package models

import (
	"fmt"
	"log"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/riteshco/Feasto/pkg/types"
)

func RegisterUser (user types.UserRegisterDB) (bool , error) {


	InsertUser := "INSERT INTO Users (username , mobile_number , email , user_role , password_hash) VALUES (? , ? , ? , ? , ?)"
	_ , err := DB.Exec(InsertUser , user.Username , user.MobileNumber , user.Email , user.UserRole , user.HashedPassword)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			log.Fatal("Duplicate entry in registration")
			return false, fmt.Errorf("user already exists")
		} else {
			log.Fatal("error inserting into the database", err)
			return false, fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("User registered successfully")
		return true, nil
	}
}
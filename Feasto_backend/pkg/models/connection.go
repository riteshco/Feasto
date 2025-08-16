package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() (*sql.DB , error) {

	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	dbPort := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true" , dbUser , dbPassword , dbHost , dbPort , database)

	var err error

	for i := 0; i < 10; i++ {
		DB, err = sql.Open("mysql", dsn)
		if err == nil {
			err = DB.Ping()
			if err == nil {
				fmt.Println("Successfully connected to Database!")
				DB.SetMaxOpenConns(75)
				DB.SetMaxIdleConns(50)
				DB.SetConnMaxLifetime(5 * time.Minute)
				return DB, nil
			}
		}
		log.Printf("Database not ready, waiting 3 seconds... Attempt %d", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("database did not become ready after several attempts: %v", err)
}

func CloseDatabase() error {
	if DB != nil {
		fmt.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}
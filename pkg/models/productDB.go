package models

import (
	"fmt"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/riteshco/Feasto/pkg/types"
)

func AddFoodDB(food types.FoodToAdd) (bool , error){
	query := "INSERT INTO Products (product_name ,isavailable ,price ,category , image_url) VALUES (?,?,?,?,?)"
	_ , err := DB.Exec(query , food.ProductName , true , food.Price , food.Category , food.ImageUrl)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate entry in registration")
			return false, fmt.Errorf("product already exists")
		} else {
			fmt.Println("error inserting into the database", err)
			return false, fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("Product added successfully")
		return true, nil
	}
}

func DeleteProductDB(productID int) error {
	query := "DELETE FROM Products WHERE id = ?"
	_ , err := DB.Exec(query , productID)
	if err != nil {
		fmt.Println("error deleting product from database:", err)
		return fmt.Errorf("error in database")
	}
	return nil
}
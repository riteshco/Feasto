package models

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/riteshco/Feasto/pkg/types"
)

func AddFoodDB(food types.FoodToAdd) (bool , int , error){
	query := "INSERT INTO Products (product_name ,isavailable ,price ,category , image_url) VALUES (?,?,?,?,?)"
	_ , err := DB.Exec(query , food.ProductName , true , food.Price , food.Category , food.ImageUrl)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate entry in registration")
			return false, http.StatusBadRequest , fmt.Errorf("product already exists")
		} else {
			fmt.Println("error inserting into the database", err)
			return false, http.StatusInternalServerError ,fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("Product added successfully")
		return true, http.StatusOK ,nil
	}
}

func UpdateFoodDB(food types.FoodToAdd, foodID int) (bool , int , error){
	query := "UPDATE Products SET product_name = ? , isavailable=? , price=? , category=? , image_url=? WHERE id = ?"
	_ , err := DB.Exec(query , food.ProductName , true , food.Price , food.Category , food.ImageUrl , foodID)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate entry in registration")
			return false, http.StatusBadRequest , fmt.Errorf("product name already exists")
		} else {
			fmt.Println("error updating into the database", err)
			return false, http.StatusInternalServerError ,fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("Product details updated successfully")
		return true, http.StatusOK ,nil
	}
}

func DeleteProductDB(productID int) (int ,error) {
	query := "DELETE FROM Products WHERE id = ?"
	_ , err := DB.Exec(query , productID)
	if err != nil {
		fmt.Println("error deleting product from database:", err)
		return http.StatusInternalServerError , fmt.Errorf("error in database")
	}
	return http.StatusOK , nil
}

func GetProductsDB() ( []types.Product , int , error) {
	query := `SELECT id , product_name , price , category , image_url FROM Products`
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	rows, err := DB.QueryContext(ctx, query)
    if err != nil {
        return nil,http.StatusInternalServerError , fmt.Errorf("error fetching products: %v", err)
    }
    defer rows.Close()

	var products = make([]types.Product, 0, 1000)

    for rows.Next() {
        var p types.Product
        if err := rows.Scan(&p.Id, &p.ProductName, &p.Price, &p.Category, &p.ImageUrl); err != nil {
            return nil, http.StatusInternalServerError, fmt.Errorf("error scanning row: %v", err)
        }
        products = append(products, p)
    }

    if err := rows.Err(); err != nil {
        return nil,http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return products , http.StatusOK, nil
}
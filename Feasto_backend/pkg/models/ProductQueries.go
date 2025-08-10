package models

import (
	"fmt"
	"net/http"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/riteshco/Feasto/pkg/types"
)

func AddFoodDB(food types.FoodToAdd) (bool , int , error){
	query := "INSERT INTO Products (product_name ,isavailable ,price ,category , image_url) VALUES (?,?,?,?,?)"
	_ , err := DB.Exec(query , food.ProductName , true , food.Price , food.Category , food.ImageUrl)
	if err != nil {
		if mysqlErr, ok := err.(*mysqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate entry in registration")
			return false, http.StatusAlreadyReported , fmt.Errorf("product already exists")
		} else {
			fmt.Println("error inserting into the database", err)
			return false, http.StatusInternalServerError ,fmt.Errorf("error in database")
		}
	} else {
		fmt.Println("Product added successfully")
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
	query := `SELECT * FROM Products`
	
	rows, err := DB.Query(query)
    if err != nil {
        return nil,http.StatusInternalServerError , fmt.Errorf("error fetching products: %v", err)
    }
    defer rows.Close()

    var orders []types.Product

    for rows.Next() {
        var p types.Product
        if err := rows.Scan(&p.Id, &p.ProductName, &p.IsAvailable, &p.Price, &p.Category, &p.ImageUrl); err != nil {
            return nil, http.StatusInternalServerError, fmt.Errorf("error scanning row: %v", err)
        }
        orders = append(orders, p)
    }

    if err := rows.Err(); err != nil {
        return nil,http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return orders , http.StatusOK, nil
}
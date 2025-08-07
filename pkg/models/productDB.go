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

func GetAllOrdersDB() ([]types.Order , error){
	query := "SELECT * FROM Orders"
	
	rows, err := DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error fetching orders: %v", err)
    }
    defer rows.Close()

    var orders []types.Order

    for rows.Next() {
        var o types.Order
        if err := rows.Scan(&o.Id, &o.CreatedAt, &o.CurrentStatus, &o.CustomerId, &o.ChefId, &o.TableNumber , &o.Instructions); err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }
        orders = append(orders, o)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }

    return orders, nil
}
package models

import (
	"fmt"

	"github.com/riteshco/Feasto/pkg/types"
)

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

func GetOrdersByCustomerId(customerID int) ([]types.Order , error){
	query := "SELECT * FROM Orders Where customer_id = ?"
	
	rows, err := DB.Query(query , customerID)
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

func InsertOrderItemDB(CustomerID int,productID int) error {
	query := "INSERT INTO OrderItems (customer_id , product_id) VALUES (? , ?)"

	_ , err := DB.Exec(query , CustomerID , productID)
	if err !=nil{
		fmt.Println("error inserting into the database", err)
		return fmt.Errorf("error in database")
	}
	return nil
}

func RemoveOrderItemDB(customerID int, ItemID int) error {
	query := "DELETE FROM OrderItems WHERE customer_id = ? AND id = ? AND order_id IS NULL"
	_, err := DB.Exec(query, customerID, ItemID)
	if err != nil {
		fmt.Println("error deleting order item from database:", err)
		return fmt.Errorf("error in database")
	}
	return nil
}

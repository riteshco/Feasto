package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/riteshco/Feasto/pkg/types"
)

func CheckIfOrderLegitDB(CustomerID int) (int , []types.OrderItem , error) {
	query := `SELECT * FROM OrderItems WHERE customer_id = ? AND order_id IS NULL`

	rows, err := DB.Query(query , CustomerID)
    if err != nil {
        return http.StatusInternalServerError, nil , fmt.Errorf("error fetching orders: %v", err)
    }
    defer rows.Close()

    var orderItems []types.OrderItem

    for rows.Next() {
        var oi types.OrderItem
        if err := rows.Scan(&oi.Id, &oi.OrderId, &oi.CustomerId, &oi.ProductId, &oi.Quantity); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound , nil ,fmt.Errorf("no Order/Cart items were found in DB with Customer ID : %d" , CustomerID)
			}
            return http.StatusInternalServerError, nil , fmt.Errorf("error scanning row: %v", err)
        }
        orderItems = append(orderItems, oi)
    }

    if err := rows.Err(); err != nil {
        return http.StatusInternalServerError, nil , fmt.Errorf("error iterating rows: %v", err)
    }
	return http.StatusOK , orderItems , nil
}

func InsertUserOrderDB(CustomerID int , TableNumber int , Instructions string) (int , int , error) {
	query := `INSERT INTO Orders (customer_id , table_number , instructions) VALUES ( ? , ? , ?)`
	result , err := DB.Exec(query , CustomerID , TableNumber , Instructions)
	if err != nil {
		fmt.Println("error registering order in database:", err)
		return http.StatusInternalServerError , 0 , fmt.Errorf("database update error")
	}
	orderIdInserted , err := result.LastInsertId()
	if err != nil {
		fmt.Println("error checking inserted rows:", err)
		return http.StatusInternalServerError , 0 , fmt.Errorf("could not verify database update")
	}

	return http.StatusOK , int(orderIdInserted) , nil
}

func UpdateOrderItemsDB(CustomerID int , OrderID int) ( int , error ){
	query := "UPDATE OrderItems SET order_id = ? WHERE customer_id = ? AND order_id IS NULL"
	result , err := DB.Exec(query , OrderID , CustomerID)
	if err != nil {
		fmt.Println("error updating orderItems in database for current order:", err)
		return http.StatusInternalServerError , fmt.Errorf("database update error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error checking affected rows:", err)
		return http.StatusInternalServerError , fmt.Errorf("could not verify database update")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound , fmt.Errorf("no orderItems found for given order ID and customer ID to ipdate")
	}

	return http.StatusOK , nil
}

func GetPricesDB(OrderID int) ([]types.Prices, int , error) {
	query := `
            SELECT Products.price 
            FROM OrderItems 
            JOIN Products ON OrderItems.product_id = Products.id 
            WHERE OrderItems.order_id = ?
            `
	rows, err := DB.Query(query , OrderID)
    if err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error fetching prices from Products: %v", err)
    }
    defer rows.Close()

    var prices []types.Prices

    for rows.Next() {
        var p types.Prices
        if err := rows.Scan(&p.Price); err != nil {
            return nil, http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        prices = append(prices, p)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return prices, http.StatusOK ,nil
}

func GetAllOrdersDB() ([]types.Order , int , error){
	query := "SELECT * FROM Orders"
	
	rows, err := DB.Query(query)
    if err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error fetching orders: %v", err)
    }
    defer rows.Close()

    var orders []types.Order

    for rows.Next() {
        var o types.Order
        if err := rows.Scan(&o.Id, &o.CreatedAt, &o.CurrentStatus, &o.CustomerId, &o.ChefId, &o.TableNumber , &o.Instructions); err != nil {
            return nil, http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        orders = append(orders, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return orders,http.StatusOK , nil
}

func DeleteOrderDB(customerID int , OrderID int) (int , error) {
	query := "DELETE FROM Orders WHERE id = ? AND customer_id = ?"
	result , err := DB.Exec(query , OrderID , customerID)
	if err != nil {
		fmt.Println("error deleting order in database for bill generation:", err)
		return http.StatusInternalServerError , fmt.Errorf("database update error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error checking affected rows:", err)
		return http.StatusInternalServerError , fmt.Errorf("could not verify database update")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound , fmt.Errorf("no order found for given order ID and customer to delete")
	}

	return http.StatusOK , nil
}

func CompleteOrderDB(OrderID int) (int , error) {
	query := `UPDATE Orders SET current_status = "delivered" WHERE id = ?`

	result , err := DB.Exec(query , OrderID)
	if err != nil {
		return http.StatusInternalServerError , fmt.Errorf("error in updating order status")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error checking affected rows:", err)
		return http.StatusInternalServerError , fmt.Errorf("could not verify database update")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound , fmt.Errorf("no order found for given order ID and customer to completing")
	}

	return http.StatusOK , nil
}

func GetOrdersByCustomerIdDB(customerID int) ([]types.Order , int , error){
	query := "SELECT * FROM Orders Where customer_id = ?"
	
	rows, err := DB.Query(query , customerID)
    if err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error fetching orders: %v", err)
    }
    defer rows.Close()

    var orders []types.Order

    for rows.Next() {
        var o types.Order
        if err := rows.Scan(&o.Id, &o.CreatedAt, &o.CurrentStatus, &o.CustomerId, &o.ChefId, &o.TableNumber , &o.Instructions); err != nil {
            return nil, http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        orders = append(orders, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return orders, http.StatusOK , nil
}

func GetCartOrderItemsDB(CustomerID int) ([]types.OrderItem , int , error) {
	query := "SELECT * FROM OrderItems WHERE customer_id = ? AND order_id IS NULL"
	rows , err := DB.Query(query , CustomerID)
	if err != nil {
        return nil , http.StatusInternalServerError , fmt.Errorf("error fetching orderItems: %v", err)
    }
    defer rows.Close()

    var orderItems []types.OrderItem

    for rows.Next() {
        var o types.OrderItem
        if err := rows.Scan(&o.Id, &o.OrderId, &o.CustomerId, &o.ProductId, &o.Quantity); err != nil {
            return nil ,  http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        orderItems = append(orderItems, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

	return orderItems , http.StatusOK , nil
}

func GetCartProductItemsDB(CustomerID int) ([]types.Product , int , error) {
	query := "SELECT id , product_name , price FROM Products WHERE id IN (SELECT product_id FROM OrderItems WHERE customer_id = ? AND order_id IS NULL)"
	rows , err := DB.Query(query , CustomerID)
	if err != nil {
        return nil , http.StatusInternalServerError , fmt.Errorf("error fetching cartProducts: %v", err)
    }
    defer rows.Close()

    var cartProducts []types.Product

    for rows.Next() {
        var o types.Product
        if err := rows.Scan(&o.Id, &o.ProductName , &o.Price); err != nil {
            return nil ,  http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        cartProducts = append(cartProducts, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

	return cartProducts , http.StatusOK , nil
}

func GetCartOrderItemsByOrderIdDB(CustomerID int , OrderID int) ([]types.OrderItem , int , error) {
	query := "SELECT * FROM OrderItems WHERE customer_id = ? AND order_id = ?"
	rows , err := DB.Query(query , CustomerID , OrderID)
	if err != nil {
        return nil , http.StatusInternalServerError , fmt.Errorf("error fetching orderItems: %v", err)
    }
    defer rows.Close()

    var orderItems []types.OrderItem

    for rows.Next() {
        var o types.OrderItem
        if err := rows.Scan(&o.Id, &o.OrderId, &o.CustomerId, &o.ProductId, &o.Quantity); err != nil {
            return nil ,  http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        orderItems = append(orderItems, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

	return orderItems , http.StatusOK , nil
}

func GetCartProductItemsByOrderIdDB(CustomerID int , OrderID int) ([]types.Product , int , error) {
	query := "SELECT id , product_name , price FROM Products WHERE id IN (SELECT product_id FROM OrderItems WHERE customer_id = ? AND order_id = ?)"
	rows , err := DB.Query(query , CustomerID , OrderID)
	if err != nil {
        return nil , http.StatusInternalServerError , fmt.Errorf("error fetching cartProducts: %v", err)
    }
    defer rows.Close()

    var cartProducts []types.Product

    for rows.Next() {
        var o types.Product
        if err := rows.Scan(&o.Id, &o.ProductName , &o.Price); err != nil {
            return nil ,  http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        cartProducts = append(cartProducts, o)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

	return cartProducts , http.StatusOK , nil
}

func InsertOrderItemsDB(CustomerID int,productID int,quantity int) (int , error) {
	query := "INSERT INTO OrderItems (customer_id , product_id , quantity) VALUES (? , ? , ?)"

	_ , err := DB.Exec(query , CustomerID , productID , quantity)
	if err !=nil{
		fmt.Println("error inserting into the database", err)
		return http.StatusInternalServerError , fmt.Errorf("error in database")
	}
	return http.StatusOK , nil
}


func RemoveOrderItemDB(customerID int, ItemID int) (int ,error) {
	query := "DELETE FROM OrderItems WHERE customer_id = ? AND id = ? AND order_id IS NULL"
	_, err := DB.Exec(query, customerID, ItemID)
	if err != nil {
		fmt.Println("error deleting order item from database:", err)
		return http.StatusInternalServerError , fmt.Errorf("error in database")
	}
	return http.StatusOK , nil
}

func AcceptOrderDB(paymentID int) (int , error) {
	query := `
		UPDATE Orders 
		SET current_status = "accepted" 
		WHERE id = (
			SELECT order_id FROM Payments WHERE id = ?
		)
	`

	result, err := DB.Exec(query, paymentID)
	if err != nil {
		fmt.Println("error accepting order in database for bill generation:", err)
		return http.StatusInternalServerError , fmt.Errorf("database update error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error checking affected rows:", err)
		return http.StatusInternalServerError , fmt.Errorf("could not verify database update")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound , fmt.Errorf("no payment found for given order ID to accept the order")
	}

	return http.StatusOK , nil
}


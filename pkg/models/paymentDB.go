package models

import (
	"fmt"
	"net/http"

	"github.com/riteshco/Feasto/pkg/types"
)

func GetAllPaymentsDB() ([]types.Payment , error){
	query := "SELECT * FROM Payments"
	
	rows, err := DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error fetching payments: %v", err)
    }
    defer rows.Close()

    var payments []types.Payment

    for rows.Next() {
        var p types.Payment
        if err := rows.Scan(&p.Id, &p.UserId, &p.OrderId, &p.TotalPayment, &p.PaymentStatus); err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }
        payments = append(payments, p)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }

    return payments, nil
}

func PaymentStatusCompleteDB(CustomerID int , PaymentID int) (int , error) {
    query := `UPDATE Payments SET payment_status = "completed" WHERE id = ? AND user_id = ?`

    result , err := DB.Exec(query , PaymentID , CustomerID )
    if err != nil {
        return http.StatusInternalServerError, fmt.Errorf("error completing the payment: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error checking affected rows:", err)
		return http.StatusInternalServerError , fmt.Errorf("could not verify database update")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound , fmt.Errorf("no payment found for given payment ID and customer to complete")
	}

	return http.StatusOK , nil
}
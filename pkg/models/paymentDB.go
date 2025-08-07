package models

import (
	"fmt"

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
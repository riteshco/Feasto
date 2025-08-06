package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/riteshco/Feasto/pkg/types"
)

func GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	// short timeout for DB operations
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	GetUser := `SELECT id, username, mobile_number, email, user_role, password_hash FROM Users WHERE email = ?`

	var user types.User
	row := DB.QueryRowContext(ctx, GetUser, email)
	err := row.Scan(&user.Id, &user.Username, &user.MobileNumber, &user.Email, &user.UserRole, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {

			return types.User{}, fmt.Errorf("user not found: %w", err)
		}
		return types.User{}, fmt.Errorf("query user by email: %w", err)
	}

	return user, nil
}

func DeleteUserDB(id int) error {

	query := "DELETE FROM Users WHERE id = ?"

	result, err := DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error deleting user: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error fetching rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no user found with ID %d", id)
    }

    return nil

}

func EditUserRoleDB(newRole string , id int) error {
	query := "UPDATE Users SET user_role = ? WHERE id = ?"

	result , err := DB.Exec(query, newRole , id)
	if err != nil {
        return fmt.Errorf("error changing user role: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error fetching rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no user found with ID %d", id)
    }

	return nil
}

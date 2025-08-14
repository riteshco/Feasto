package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/riteshco/Feasto/pkg/types"
)

func GetUserByEmailDB(ctx context.Context, email string) (types.User, int , error) {
	// short timeout for DB operations
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	GetUser := `SELECT id, username, mobile_number, email, user_role, password_hash FROM Users WHERE email = ?`

	var user types.User
	row := DB.QueryRowContext(ctx, GetUser, email)
	err := row.Scan(&user.Id, &user.Username, &user.MobileNumber, &user.Email, &user.UserRole, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {

			return types.User{}, http.StatusNotFound , fmt.Errorf("user not found")
		}
		return types.User{}, http.StatusInternalServerError , fmt.Errorf("error in fetching user from DB")
	}

	return user, http.StatusOK , nil
}

func DeleteUserDB(id int) (int , error ) {

	query := "DELETE FROM Users WHERE id = ?"

	result, err := DB.Exec(query, id)
    if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error deleting user: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error fetching rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return http.StatusNotFound , fmt.Errorf("no user found with ID %d", id)
    }

    return http.StatusOK , nil

}

func EditUserRoleDB(newRole string , id int) (int , error) {
	query := "UPDATE Users SET user_role = ? WHERE id = ?"

	result , err := DB.Exec(query, newRole , id)
	if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error changing user role: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error fetching rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return http.StatusNotFound , fmt.Errorf("no user found with ID %d", id)
    }

	return http.StatusOK , nil
}

func GetAllUsersDB() ([]types.User , int , error) {
	query := "SELECT * FROM Users"

	rows, err := DB.Query(query)
    if err != nil {
		return nil , http.StatusInternalServerError , fmt.Errorf("error fetching users: %v", err)
    }
    defer rows.Close()
	
    var users []types.User
	
    for rows.Next() {
		var u types.User
        if err := rows.Scan(&u.Id, &u.Username, &u.MobileNumber, &u.Email, &u.UserRole, &u.HashedPassword, &u.ChangeRequest); err != nil {
			return nil, http.StatusInternalServerError , fmt.Errorf("error scanning row: %v", err)
        }
        users = append(users, u)
    }

    if err := rows.Err(); err != nil {
        return nil, http.StatusInternalServerError , fmt.Errorf("error iterating rows: %v", err)
    }

    return users, http.StatusOK ,nil
}

func GetSingleUserDB(id int) (types.User , int , error) {
	query := "SELECT * FROM Users Where id = ?"

	var u types.User
	row := DB.QueryRow(query, id)
	err := row.Scan(&u.Id, &u.Username, &u.MobileNumber, &u.Email, &u.UserRole, &u.HashedPassword, &u.ChangeRequest)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.User{}, http.StatusNotFound ,fmt.Errorf("user not found with ID : %d" , id)
		}
		return types.User{}, http.StatusInternalServerError , fmt.Errorf("query single product: %w", err)
	}

	return u, http.StatusOK , nil
}

func AddChangeRoleToDB(id int , newRole string) (int , error) {
	query := "UPDATE Users SET change_role_to = ? WHERE id = ?"

	result , err := DB.Exec(query, newRole , id)
	if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error changing user role request: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return http.StatusInternalServerError , fmt.Errorf("error fetching rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return http.StatusNotFound , fmt.Errorf("no user found with ID %d", id)
    }

	return http.StatusOK , nil
}
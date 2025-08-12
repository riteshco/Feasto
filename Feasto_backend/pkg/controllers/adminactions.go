package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/riteshco/Feasto/pkg/constants"
	"github.com/riteshco/Feasto/pkg/types"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/models"
)


func DeleteUserAPI(w http.ResponseWriter , r *http.Request){
	vars:= mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	http.Error(w, "Invalid user ID", http.StatusBadRequest)
    	return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		status , err := models.DeleteUserDB(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  err.Error(),
			})
		} else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  "User deleted successfully",
			})
		}
	} else {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
}

func EditUserRoleAPI(w http.ResponseWriter , r *http.Request){
	vars:= mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	http.Error(w, "Invalid user ID", http.StatusBadRequest)
    	return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin{
		var user_role types.UserRole

		if err := json.NewDecoder(r.Body).Decode(&user_role); err != nil {
        	http.Error(w, "Invalid JSON", http.StatusBadRequest)
    	    return
    	}
		new_role := user_role.Role
		status , err := models.EditUserRoleDB(new_role , id)
		if err != nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  err.Error(),
			})
		} else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  "User role changed successfully",
			})
		}
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized);
		return
	}
}

func GetAllUsersAPI(w http.ResponseWriter , r *http.Request) {
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		users , status , err := models.GetAllUsersDB()
		if err != nil {
        	http.Error(w, err.Error(), status)
        	return
    	}
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(users)
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized)
		return
	}
}

func GetSingleUserAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	http.Error(w, "Invalid user ID", http.StatusBadRequest)
    	return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		user , status , err := models.GetSingleUserDB(id)
		if err != nil {
        	http.Error(w, err.Error(), status)
        	return
    	}
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized)
		return
	}
}

func GetAllOrdersAPI(w http.ResponseWriter , r *http.Request) {
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin || UserRole == constants.RoleChef {
		orders, status , err := models.GetAllOrdersDB()
		if err != nil {
        	http.Error(w, err.Error(), status)
			fmt.Println("Error in getting all-orders from DB : " , err)
        	return
    	}
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(orders)
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized)
		return
	}
}

func GetAllPaymentsAPI(w http.ResponseWriter , r *http.Request) {
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		payments, status , err := models.GetAllPaymentsDB()
		if err != nil {
        	http.Error(w, err.Error(), status)
        	return
    	}
    	w.Header().Set("Content-Type", "application/json")
    	json.NewEncoder(w).Encode(payments)
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized)
		return
	}
}

func DeleteProductAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId , err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
    	return
	} 
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		status , err := models.DeleteProductDB(productId)
		if err != nil {
			http.Error(w , "Server Error" , status)
			return
		} else {
			w.WriteHeader(status)
			w.Write([]byte("Product Deleted Successfully!!"))
		}
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}
}

func GenBillAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	orderId , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Order Id" , http.StatusBadRequest)
		return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleAdmin {
		statusCode , err := models.AcceptOrderDB(orderId)
		if err != nil {
			http.Error(w , err.Error() , statusCode)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Bill generated Successfully!!"))
		}
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}
}
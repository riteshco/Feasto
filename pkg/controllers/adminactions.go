package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/riteshco/Feasto/pkg/types"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/models"
)


func DeleteUser(w http.ResponseWriter , r *http.Request){
	vars:= mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	http.Error(w, "Invalid user ID", http.StatusBadRequest)
    	return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == "admin" {
		err := models.DeleteUserDB(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  err.Error(),
			})
		} else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  "User deleted successfully",
			})
		}
	} else {
		http.Error(w, "unauthorized", http.StatusUnauthorized); return
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
	if UserRole == "admin" {
		var user_role types.UserRole

		if err := json.NewDecoder(r.Body).Decode(&user_role); err != nil {
        	http.Error(w, "Invalid JSON", http.StatusBadRequest)
    	    return
    	}
		new_role := user_role.Role
		err := models.EditUserRoleDB(new_role , id)
		if err != nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  err.Error(),
			})
		} else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  "User role changed successfully",
			})
		}
	} else {
		http.Error(w, "unauthorized access", http.StatusUnauthorized); return
	}
}
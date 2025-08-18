package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/constants"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/utils"
)

func OrderDoneAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return
	}
	ChefID := r.Context().Value("id").(int)
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleChef {
		status , err := models.CompleteOrderDB(OrderId , ChefID)
		if err != nil {
			utils.ErrorHandling(w , err.Error() , status)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Completed Successfully!!"))
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}
}

func TakeOrderAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return
	}
	ChefID := r.Context().Value("id").(int)
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleChef {
		status , err := models.TakeOrderDB(OrderId , ChefID)
		if err != nil {
			utils.ErrorHandling(w , err.Error() , status)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Completed Successfully!!"))
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}
}

func DeliveredOrdersAPI(w http.ResponseWriter , r *http.Request){
	ChefID := r.Context().Value("id").(int)
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleChef {
	orders ,status , err := models.GetDeliveredOrdersByChefIdDB(ChefID)
	if err != nil {
		utils.ErrorHandling(w , err.Error() , status)
		fmt.Println("Error in getting orders from Database : " , err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}

}
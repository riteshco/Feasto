package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/models"
)

func OrderDoneAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return
	}
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == "chef" {
		status , err := models.CompleteOrderDB(OrderId)
		if err != nil {
			http.Error(w , err.Error() , status)
			return
		}
		http.Error(w , "Order Completed Successfully!" , status)
	} else {
		http.Error(w , "Unauthorized access!" , http.StatusUnauthorized)
		return
	}
}
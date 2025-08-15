package controllers

import (
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
	UserRole := r.Context().Value("user_role").(string)
	if UserRole == constants.RoleChef {
		status , err := models.CompleteOrderDB(OrderId)
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
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
)

func UserOrders(w http.ResponseWriter , r *http.Request){
	CustomerID := r.Context().Value("id").(int)
	orders , err := models.GetOrdersByCustomerId(CustomerID)
	if err != nil {
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		fmt.Println("Error in getting orders from Database : " , err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)

}

func AddOneToCartAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId, err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
    	http.Error(w, "Invalid product ID", http.StatusBadRequest)
    	return
	}
	err = models.InsertOrderItemDB(CustomerID , productId)
	if err != nil {
		http.Error(w , "Server Error" , http.StatusInternalServerError)
		fmt.Println("Error in inserting orderItem by ID in DB : " , err)
		return
	} else {
		http.Error(w , "Added to Cart successfully!!" , http.StatusOK)
	}
}

func RemoveFromCartAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	ItemId, err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Item ID", http.StatusBadRequest)
		return
	}
	err = models.RemoveOrderItemDB(CustomerID , ItemId)
	if err != nil {
		http.Error(w , "Server Error" , http.StatusInternalServerError)
		fmt.Println("Error in removing OrderItem in DB : " , err)
		return
	} else {
		http.Error(w , "Removed from Cart Successfully!!" , http.StatusOK)
	}
}
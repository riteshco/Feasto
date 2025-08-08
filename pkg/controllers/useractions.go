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

func AddToCartAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId , err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	CustomerID := r.Context().Value("id").(int)
	quantStr := vars["quantity"]
	quantity , err := strconv.Atoi(quantStr)
	if err != nil || quantity <= 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}
	err = models.InsertOrderItemsDB(CustomerID , productId , quantity)
	if err != nil {
		http.Error(w , "Server Error" , http.StatusInternalServerError)
		fmt.Println("Error in inserting orderItem by ID and quantity in DB : " , err)
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

func DeleteOrderAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
	}
	status , err := models.DeleteOrderDB(CustomerID , OrderId)
	if err != nil {
		fmt.Println("Error in deleting order in DB : " , err)
		http.Error(w , err.Error() , status)
		return
	} else {
		http.Error(w , "Deleted Order Successfully!!" , status)
	}
}

func PaymentDoneAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	PaymentId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Payment ID" , http.StatusBadRequest)
	}
	status , err := models.PaymentStatusCompleteDB(CustomerID , PaymentId)
	if err != nil {
		fmt.Println("Error in completing payment : " , err)
		http.Error(w , err.Error() , status)
		return
	} else {
		http.Error(w , "Payment Completed Successfully!!" , status)
	}
}

func GetPaymentThroughOrderAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
	}
	payment , status , err := models.GetPaymentThroughOrderDB(OrderId , CustomerID)
	if err != nil {
		fmt.Println("Error in getting payment through order id : " , err)
		http.Error(w , err.Error() , status)
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payment)

}

func CartOrderAPI(w http.ResponseWriter , r *http.Request) {
	CustomerID := r.Context().Value("id").(int)

	status , orderItems , err := models.CheckIfOrderLegit(CustomerID)
	if err != nil {
		fmt.Println("Error in ordering : " , err)
		http.Error(w , err.Error() , status)
		return
	}

	var order types.RegisterOrder

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	TableNumber := order.TableNumber
	Instructions := order.Instructions

	status , OrderId , err := models.InsertUserOrderDB(CustomerID , TableNumber , Instructions)
	if err != nil {
		fmt.Println("Error in ordering : " , err)
		http.Error(w , err.Error() , status)
		return
	} else {
		status , err := models.UpdateOrderItemsDB(CustomerID , OrderId)
		if err != nil {
			fmt.Println("Error in ordering : " , err)
			http.Error(w , err.Error() , status)
			return
		}
		prices , err := models.GetPricesDB(OrderId)
		if err != nil {
			fmt.Println("Error in getting prices from DB : " , err)
			http.Error(w , err.Error() , http.StatusInternalServerError)
			return
		}
		var totalAmount float64
        for i := range orderItems {
    	productPrice := prices[i].Price
    	quantity := orderItems[i].Quantity
    	totalAmount += productPrice * float64(quantity)
		}
		status , err = models.InsertPaymentDB(CustomerID , OrderId , totalAmount)
		if err != nil {
			fmt.Println("Error in inserting payment in db : ", err)
			http.Error(w , err.Error() , status)
			return
		}
		http.Error(w , "Order registered Successfully!!" , status)
	}
	
}
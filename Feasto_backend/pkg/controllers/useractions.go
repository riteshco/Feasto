package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
	"github.com/riteshco/Feasto/pkg/utils"
)

func UserOrdersAPI(w http.ResponseWriter , r *http.Request){
	CustomerID := r.Context().Value("id").(int)
	orders ,status , err := models.GetOrdersByCustomerIdDB(CustomerID)
	if err != nil {
		utils.ErrorHandling(w , err.Error() , status)
		fmt.Println("Error in getting orders from Database : " , err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)

}

func UserPastOrdersAPI(w http.ResponseWriter , r *http.Request){
	CustomerID := r.Context().Value("id").(int)
	orders ,status , err := models.GetPastOrdersByCustomerIdDB(CustomerID)
	if err != nil {
		utils.ErrorHandling(w , err.Error() , status)
		fmt.Println("Error in getting orders from Database : " , err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)

}


func AddToCartAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId , err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorHandling(w , "Invalid quantity" , http.StatusBadRequest)
		return
	}
	CustomerID := r.Context().Value("id").(int)
	quantStr := vars["quantity"]
	quantity , err := strconv.Atoi(quantStr)
	if err != nil || quantity <= 0 {
		utils.ErrorHandling(w, "Invalid quantity", http.StatusBadRequest)
		return
	}
	status , err := models.InsertOrderItemsDB(CustomerID , productId , quantity)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", http.StatusBadRequest)
		fmt.Println("Error in inserting orderItem by ID and quantity in DB : " , err)
		return
	} else {
		w.WriteHeader(status)
		w.Write([]byte("Added to Cart Successfully!!"))
	}
}

func GetCartItemsAPI(w http.ResponseWriter , r *http.Request){
	CustomerID := r.Context().Value("id").(int)
	orderItems , status , err := models.GetCartOrderItemsDB(CustomerID)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in getting OrderItem in DB : " , err)
		return
	}
	productItems , status , err := models.GetCartProductItemsDB(CustomerID)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in getting ProductItems in DB : " , err)
		return
	}
	response := struct {
        Orders   interface{} `json:"orders"`
        Products interface{} `json:"products"`
    }{
        Orders:   orderItems,
        Products: productItems,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func OrderItemsAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	IdStr := vars["id"]
	OrderID , err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return 
	}
	CustomerID := r.Context().Value("id").(int)

	orderItems , status , err := models.GetCartOrderItemsByOrderIdDB(CustomerID , OrderID)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in getting OrderItem by OrderID in DB : " , err)
		return
	}
	productItems , status , err := models.GetCartProductItemsByOrderIdDB(CustomerID , OrderID)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in getting ProductItems by OrderID in DB : " , err)
		return
	}
	response := struct {
        Orders   interface{} `json:"orders"`
        Products interface{} `json:"products"`
    }{
        Orders:   orderItems,
        Products: productItems,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
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
	status , err := models.RemoveOrderItemDB(CustomerID , ItemId)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in removing OrderItem in DB : " , err)
		return
	} else {
		w.WriteHeader(status)
		w.Write([]byte("Removed from Cart Successfully!!"))
	}
}

func AddChangeRequestAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	role := vars["role"]
	CustomerID := r.Context().Value("id").(int)
	status , err := models.AddChangeRoleToDB(CustomerID , role)
	if err != nil {
		utils.ErrorHandling(w, "Server Error", status)
		fmt.Println("Error in requesting role change : " , err)
		return
	} else {
		w.WriteHeader(status)
		w.Write([]byte("Added change request Successfully!!"))
	}
}

func DeleteOrderAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return
	}
	status , err := models.DeleteOrderDB(CustomerID , OrderId)
	if err != nil {
		fmt.Println("Error in deleting order in DB : " , err)
		utils.ErrorHandling(w, err.Error(), status)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted Order Successfully!!"))
	}
}

func PaymentDoneAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	PaymentId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Payment ID" , http.StatusBadRequest)
		return
	}
	status , err := models.PaymentStatusCompleteDB(CustomerID , PaymentId)
	if err != nil {
		fmt.Println("Error in completing payment : " , err)
		utils.ErrorHandling(w, err.Error(), status)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Payment Completed Successfully!!"))
	}
}

func GetPaymentThroughOrderAPI(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	OrderId , err := strconv.Atoi(idStr)
	CustomerID := r.Context().Value("id").(int)
	if err != nil {
		http.Error(w , "Invalid Order ID" , http.StatusBadRequest)
		return
	}
	payment , status , err := models.GetPaymentThroughOrderDB(OrderId , CustomerID)
	if err != nil {
		fmt.Println("Error in getting payment through order id : " , err)
		utils.ErrorHandling(w, err.Error(), status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payment)

}

func CartOrderAPI(w http.ResponseWriter , r *http.Request) {
	CustomerID := r.Context().Value("id").(int)

	status , orderItems , err := models.CheckIfOrderLegitDB(CustomerID)
	if err != nil {
		fmt.Println("Error in ordering : " , err)
		utils.ErrorHandling(w, err.Error(), status)
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
		utils.ErrorHandling(w, err.Error(), status)
		return
	} else {
		status , err := models.UpdateOrderItemsDB(CustomerID , OrderId)
		if err != nil {
			fmt.Println("Error in ordering : " , err)
			utils.ErrorHandling(w, err.Error(), status)
			return
		}
		prices, status , err := models.GetPricesDB(OrderId)
		if err != nil {
			fmt.Println("Error in getting prices from DB : " , err)
			utils.ErrorHandling(w, err.Error(), status)
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
			utils.ErrorHandling(w, err.Error(), status)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Registered Successfully!!"))
	}
	
}
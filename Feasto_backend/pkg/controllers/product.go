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

func AddFoodAPI(w http.ResponseWriter , r *http.Request){
	var food types.FoodToAdd
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	
	if food.ProductName == "" || food.Category == "" || food.Price <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Message{Message: "All fields (ProductName, Category, Price) are required and Price must be greater than 0"})
		return
	}

	success , status , err := models.AddFoodDB(food)
	if err != nil {
		fmt.Println("Could not log product")
		utils.ErrorHandling(w , err.Error() , status)
		return
	}
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product Added Successfully!!"))
		return
	}

}

func UpdateFoodAPI(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId , err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorHandling(w , "Invalid ID" , http.StatusBadRequest)
		return
	}
	var food types.FoodToAdd
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	
	if food.ProductName == "" || food.Category == "" || food.Price <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.Message{Message: "All fields (ProductName, Category, Price) are required and Price must be greater than 0"})
		return
	}

	success , status , err := models.UpdateFoodDB(food , productId)
	if err != nil {
		fmt.Println("Could not log product")
		utils.ErrorHandling(w , err.Error() , status)
		return
	}
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product details updated Successfully!!"))
		return
	}

}

func GetAllProductsAPI(w http.ResponseWriter , r *http.Request) {
	products, status , err := models.GetProductsDB()
	if err != nil {
		fmt.Println("Error in getting the products from DB : " , err)
		utils.ErrorHandling(w , err.Error() , status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
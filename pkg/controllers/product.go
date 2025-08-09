package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riteshco/Feasto/pkg/models"
	"github.com/riteshco/Feasto/pkg/types"
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
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), status)
		return
	}
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product Added Successfully!!"))
		return
	}

}

func GetAllProductsAPI(w http.ResponseWriter , r *http.Request) {
	products, status , err := models.GetProductsDB()
	if err != nil {
		fmt.Println("Error in getting the products from DB : " , err)
		http.Error(w , err.Error() , status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
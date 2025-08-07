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
	
	if food.ProductName == "" || food.Category =="" || food.Price <= 0 {
		fmt.Println("All fields are required to register!")
		toSend := types.Message{Message: "All fields are required to register!"}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusBadRequest)
		return
	}

	success , err := models.AddFoodDB(food)
	if err != nil {
		fmt.Println("Could not log product")
		toSend := types.Message{Message: err.Error()}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusInternalServerError)
		return
	}
	if success {
		fmt.Println("Product added successfully")
		toSend := types.Message{Message: "Product added successfully"}
		b, err := json.Marshal(toSend)
		if err != nil {
			fmt.Println(err, "could not marshal message")
		}
		http.Error(w, string(b), http.StatusOK)
		return
	}

}
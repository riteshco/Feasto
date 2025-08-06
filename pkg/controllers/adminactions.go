package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func DeleteUser(w http.ResponseWriter , r *http.Request){
	vars:= mux.Vars(r)
	fmt.Println(vars)
}
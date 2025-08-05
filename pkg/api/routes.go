package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/controllers"
)


func Run(){
	router := mux.NewRouter()


	router.HandleFunc("/" , test_handler ).Methods("GET")

	//--Registration--
	//router.HandleFunc("/register" , controllers.RegisterPage).Methods("GET")
	router.HandleFunc("/register" , controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/register" , controllers.RegisterAPIUser).Methods("POST")

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000" , router)
}

func test_handler(w http.ResponseWriter , r * http.Request){
	fmt.Fprintf(w , "Hello World!")
}
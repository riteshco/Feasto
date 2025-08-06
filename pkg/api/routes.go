package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/controllers"
	"github.com/riteshco/Feasto/pkg/middleware"
)


func Run(){
	router := mux.NewRouter()


	router.HandleFunc("/" , test_handler ).Methods("GET")

	//--Registration--
	//router.HandleFunc("/register" , controllers.RegisterPage).Methods("GET")
	router.HandleFunc("/register" , controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/register" , controllers.RegisterAPIUser).Methods("POST")

	//--authentication--
	// router.HandleFunc("/login" , controllers.LoginPage).Methods("GET")
	router.HandleFunc("/auth" , controllers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/auth" , controllers.AuthenticateUserAPI).Methods("POST")

	//--Pages--
	// router.HandleFunc("/home", JWTAuthMiddleware(http.HandleFunc(controllers.Homepage))).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)
	api.HandleFunc("/delete-user/{id}" , controllers.DeleteUser).Methods("POST")


	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000" , router)
}

func test_handler(w http.ResponseWriter , r * http.Request){
	fmt.Fprintf(w , "Hello World!")
}
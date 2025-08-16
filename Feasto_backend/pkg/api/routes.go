package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riteshco/Feasto/pkg/controllers"
	"github.com/riteshco/Feasto/pkg/middleware"
	"github.com/rs/cors"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()


	router.HandleFunc("/" , test_handler ).Methods("GET")

	//--Registration--
	router.HandleFunc("/api/register" , controllers.RegisterUserAPI).Methods("POST")

	//--authentication--
	router.HandleFunc("/api/auth" , controllers.AuthenticateUserAPI).Methods("POST")

	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.JWTAuthMiddleware)
	

	authRouter.HandleFunc("/past-orders", controllers.UserPastOrdersAPI).Methods("GET")

	authRouter.HandleFunc("/order/items/{id:[0-9]+}", controllers.OrderItemsAPI).Methods("GET")

	authRouter.HandleFunc("/order/payment/{id:[0-9]+}", controllers.GetPaymentThroughOrderAPI).Methods("GET")
	
	authRouter.HandleFunc("/all-products" , controllers.GetAllProductsAPI).Methods("GET")
	
	authRouter.HandleFunc("/add-food" , controllers.AddFoodAPI).Methods("POST")
	
	// -- Admin Actions --
	authRouter.HandleFunc("/users/{id:[0-9]+}" , controllers.GetSingleUserAPI).Methods("GET")
	
	authRouter.HandleFunc("/all-payments" , controllers.GetAllPaymentsAPI).Methods("GET")
	
	authRouter.HandleFunc("/all-orders", controllers.GetAllOrdersAPI).Methods("GET")
	
	authRouter.HandleFunc("/users" , controllers.GetAllUsersAPI).Methods("GET")
	
	authRouter.HandleFunc("/delete-product/{id:[0-9]+}", controllers.DeleteProductAPI).Methods("DELETE")
	
	authRouter.HandleFunc("/gen-bill/{id:[0-9]+}" , controllers.GenBillAPI).Methods("POST")
	
	// -- User actions --
	authRouter.HandleFunc("/orders", controllers.UserOrdersAPI).Methods("GET")

	authRouter.HandleFunc("/add-to-cart/{id:[0-9]+}/{quantity:[0-9]+}" , controllers.AddToCartAPI).Methods("POST")
	
	authRouter.HandleFunc("/remove-from-cart/{id:[0-9]+}" , controllers.RemoveFromCartAPI).Methods("POST")
	
	authRouter.HandleFunc("/delete-order/{id:[0-9]+}" , controllers.DeleteOrderAPI).Methods("DELETE")
	
	authRouter.HandleFunc("/payment-done/{id:[0-9]+}" , controllers.PaymentDoneAPI).Methods("POST")
	
	authRouter.HandleFunc("/cart/order", controllers.CartOrderAPI).Methods("POST")

	authRouter.HandleFunc("/cartItems", controllers.GetCartItemsAPI).Methods("GET")

	authRouter.HandleFunc("/change_role_request/{role}", controllers.AddChangeRequestAPI).Methods("POST")

	// -- Chef specific action --
	authRouter.HandleFunc("/order-done/{id:[0-9]+}" , controllers.OrderDoneAPI).Methods("POST")
	
	
	authRouter.HandleFunc("/edit-user-role/{id:[0-9]+}", controllers.EditUserRoleAPI).Methods("PATCH")
	
	
	authRouter.HandleFunc("/delete-user/{id:[0-9]+}" , controllers.DeleteUserAPI).Methods("DELETE")

	return router
}

func Run() {
	router := SetupRouter()
	    // Configure CORS options
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
    })

    // Wrap the router with the CORS handler
    handler := c.Handler(router)

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", handler)
}

func test_handler(w http.ResponseWriter , r * http.Request){
	fmt.Fprintf(w , "Hello World!")
}
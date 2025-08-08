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
	router.HandleFunc("/register" , controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/register" , controllers.RegisterAPIUser).Methods("POST")

	//--authentication--
	router.HandleFunc("/auth" , controllers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/auth" , controllers.AuthenticateUserAPI).Methods("POST")

	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.JWTAuthMiddleware)
	
	//--Pages (For front end)--
	// authRouter.HandleFunc("/home", controllers.Homepage).Methods("GET")
	// authRouter.HandleFunc("/add-food", controllers.AddFoodpage).Methods("GET")
	// authRouter.HandleFunc("/admin", controllers.AdminPage).Methods("GET")
	// authRouter.HandleFunc("/cart", controllers.CartPage).Methods("GET")
	// authRouter.HandleFunc("/categories", controllers.CategoriesPage).Methods("GET")
	// authRouter.HandleFunc("/past-orders", controllers.PastOrdersPage).Methods("GET")
	// authRouter.HandleFunc("/order/bill/{id:[0-9]+}", controllers.OrderBillPage).Methods("GET")
	// authRouter.HandleFunc("/order/items/{id:[0-9]+}", controllers.OrderItemsPage).Methods("GET")

	authRouter.HandleFunc("/order/payment/{id:[0-9]+}", controllers.GetPaymentThroughOrderAPI).Methods("GET")
	
	authRouter.HandleFunc("/api/all-products" , controllers.GetAllProducts).Methods("GET")
	
	authRouter.HandleFunc("/api/add-food" , controllers.AddFoodAPI).Methods("POST")
	
	// -- Admin Actions --
	authRouter.HandleFunc("/api/users/{id:[0-9]+}" , controllers.GetSingleUser).Methods("GET")
	
	authRouter.HandleFunc("/api/all-payments" , controllers.GetAllPayments).Methods("GET")
	
	authRouter.HandleFunc("/api/all-orders", controllers.GetAllOrders).Methods("GET")
	
	authRouter.HandleFunc("/api/users" , controllers.GetAllUsers).Methods("GET")
	
	authRouter.HandleFunc("/api/delete-product/{id:[0-9]+}", controllers.DeleteProductAPI).Methods("POST")
	
	authRouter.HandleFunc("/api/gen-bill/{id:[0-9]+}" , controllers.GenBillAPI).Methods("POST")
	
	// -- User actions --
	authRouter.HandleFunc("/api/orders", controllers.UserOrders).Methods("GET")
	
	authRouter.HandleFunc("/api/add-one-to-cart/{id:[0-9]+}" , controllers.AddOneToCartAPI).Methods("POST")

	authRouter.HandleFunc("/api/add-to-cart/{id:[0-9]+}/{quantity:[0-9]+}" , controllers.AddToCartAPI).Methods("POST")
	
	authRouter.HandleFunc("/api/remove-from-cart/{id:[0-9]+}" , controllers.RemoveFromCartAPI).Methods("POST")
	
	authRouter.HandleFunc("/api/delete-order/{id:[0-9]+}" , controllers.DeleteOrderAPI).Methods("POST")
	
	authRouter.HandleFunc("/api/payment-done/{id:[0-9]+}" , controllers.PaymentDoneAPI).Methods("POST")

	// -- Chef specific action --
	authRouter.HandleFunc("/api/order-done/{id:[0-9]+}" , controllers.OrderDoneAPI).Methods("POST")
	
	// ---Needs Front-end---
	// authRouter.HandleFunc("/cart/order", controllers.CartOrder).Methods("POST")
	// authRouter.HandleFunc("/new-name/{id:[0-9]+}", controllers.NewProductName).Methods("POST")
	// authRouter.HandleFunc("/new-price/{id:[0-9]+}", controllers.NewProductPrice).Methods("POST")
	// ---------------------
	
	
	authRouter.HandleFunc("/api/edit-user-role/{id:[0-9]+}", controllers.EditUserRoleAPI).Methods("POST")
	
	
	authRouter.HandleFunc("/api/delete-user/{id:[0-9]+}" , controllers.DeleteUser).Methods("POST")
	

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000" , router)
}

func test_handler(w http.ResponseWriter , r * http.Request){
	fmt.Fprintf(w , "Hello World!")
}
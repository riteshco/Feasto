package types

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
)

type Quantity struct {
	Count int `json:"quantity"`
}

type UserRole struct {
	Role string `json:"user_role"`
}

type MyClaims struct {
	ID       int `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    UserRole string `json:"user_role"`
    jwt.RegisteredClaims
}

type UserRegister struct {
	Username string `json:"username"`
	MobileNumber string `json:"mobile_number"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterDB struct {
	Username string `json:"username"`
	MobileNumber string `json:"mobile_number"`
	Email string `json:"email"`
	UserRole string `json:"user_role"`
	HashedPassword string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	MobileNumber string `json:"mobile_number"`
	Email string `json:"email"`
	UserRole string `json:"user_role"`
	HashedPassword string `json:"password"`
	ChangeRequest *string `json:"change_role_to"`
}

type Order struct {
	Id int `json:"id"`
	CreatedAt string `json:"created_at"`
	CurrentStatus string `json:"current_status"`
	CustomerId int `json:"customer_id"`
	ChefId sql.NullInt64 `json:"chef_id"`
	TableNumber int `json:"table_number"`
	Instructions string `json:"instructions"`
}

type OrderItem struct {
	Id int `json:"id"`
	OrderId sql.NullInt64 `json:"order_id"`
	CustomerId int `json:"customer_id"`
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type Prices struct {
	Price float64 `json:"price"`
}

type Product struct {
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	IsAvailable bool `json:"isavailable"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	ImageUrl sql.NullString `json:"image_url"`
}

type CartItem struct {
	Id int `json:"id"`
	ProductName string `json:"product_name"`
	Quantity int `json:"quantity"`
}

type FoodToAdd struct {
	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	ImageUrl string `json:"image_url"`
}

type Payment struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	OrderId int `json:"order_id"`
	TotalPayment float64 `json:"Total_payment"`
	PaymentStatus string `json:"payment_status"`
}

type Message struct {
	Message string `json:"message"`
}

type RegisterOrder struct {
	TableNumber int `json:"table_number"`
	Instructions string `json:"instructions"`
}
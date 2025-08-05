package types

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

type Message struct {
	Message string `json:"message"`
}
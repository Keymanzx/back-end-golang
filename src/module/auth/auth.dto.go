package auth

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string     `json:"token"`
}

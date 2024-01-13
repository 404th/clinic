package model

type User struct {
	ID        string  `json:"id" binding:"required"`
	RoleID    string  `json:"role_id" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Firstname string  `json:"firstname"`
	Surname   string  `json:"surname"`
	Wallet    float64 `json:"wallet"`
	Active    int     `json:"active"`
	Email     string  `json:"email"`
	Password  string  `json:"password" binding:"required"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CreateUserRequest struct {
	RoleID    string `json:"role_id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TransferMoneyRequest struct {
	ID    string  `json:"id" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

type GetAllUsersRequest struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}

type GetAllUsersResponse struct {
	Metadata Metadata `json:"metadata"`
	Users    []User   `json:"users"`
}

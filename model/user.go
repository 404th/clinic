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
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}

type UpdateUserRequest struct {
	ID        string  `json:"id" binding:"required"`
	RoleID    string  `json:"role_id" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Firstname string  `json:"firstname"`
	Surname   string  `json:"surname"`
	Wallet    float64 `json:"wallet"`
	Active    int     `json:"active"`
	Email     string  `json:"email"`
	Password  string  `json:"password" binding:"required"`
}

type TransferMoneyRequest struct {
	ID    string  `json:"id" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

// "id" UUID DEFAULT (uuid_generate_v4()),
// "role_id" UUID,
// "username" VARCHAR(255) UNIQUE NOT NULL,
// "firstname" VARCHAR(255) NOT NULL,
// "surname" VARCHAR(255) NOT NULL,
// "wallet" BIGSERIAL NOT NULL DEFAULT 0,
// "active" INTEGER NOT NULL DEFAULT 1,
// "password" TEXT NOT NULL,
// "created_at" TIMESTAMP DEFAULT (NOW()),
// "updated_at" TIMESTAMP DEFAULT (NOW()),
// "deleted_at" TIMESTAMP,
// PRIMARY KEY ("id", "role_id")

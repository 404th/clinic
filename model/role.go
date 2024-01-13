package model

type Role struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Rolename  string  `json:"role_name"`
	Active    int     `json:"active"`
	CreateAt  string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type CreateRoleRequest struct {
	Rolename string  `json:"rolename"`
	Price    float64 `json:"price"`
}

type GetAllRolesRequest struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}

type GetAllRolesResponse struct {
	Metadata Metadata `json:"metadata"`
	Roles    []Role   `json:"roles"`
}

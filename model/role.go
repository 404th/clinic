package model

type Role struct {
	ID        string `json:"id"`
	Rolename  string `json:"role_name"`
	Active    int    `json:"active"`
	CreateAt  string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateRoleRequest struct {
	Rolename string `json:"rolename"`
}

type ToggleActive struct {
	ID     string `json:"id"`
	Active int    `json:"active"`
}

package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	role_table_name string = "roles"
)

type role struct {
	db *pgxpool.Pool
}

func NewRole(db *pgxpool.Pool) storage.RoleI {
	return &role{
		db: db,
	}
}

func (r *role) CreateRole(ctx context.Context, req *model.CreateRoleRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	query := fmt.Sprintf(`
		INSERT INTO %s (
			rolename 
		) VALUES (
			$1
		) RETURNING id
	`, role_table_name)

	var (
		id sql.NullString
	)

	if err = r.db.QueryRow(ctx, query, req.Rolename).Scan(&id); err != nil {
		return resp, err
	}

	if id.Valid {
		resp.ID = id.String
	}

	return resp, err
}

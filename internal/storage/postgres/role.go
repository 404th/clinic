package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/jackc/pgx/v4/pgxpool"
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
			rolename,
			price  
		) VALUES (
			$1,
			$2
		) RETURNING id
	`, roles_table_name)

	var (
		id sql.NullString
	)

	if err = r.db.QueryRow(ctx, query, req.Rolename, req.Price).Scan(&id); err != nil {
		return resp, err
	}

	if id.Valid {
		resp.ID = id.String
	}

	return resp, err
}

func (r *role) GetAllRoles(ctx context.Context, req *model.GetAllRolesRequest) (resp *model.GetAllRolesResponse, err error) {
	resp = &model.GetAllRolesResponse{}

	var (
		offset int32
	)

	offset = (req.Page - 1) * req.Limit

	query := fmt.Sprintf(`
		SELECT 
			id,
			rolename,
			active,
			price 
		FROM 
			%s 
		WHERE deleted_at IS NULL 
		ORDER BY updated_at DESC 
		LIMIT $1 OFFSET $2 
	`, roles_table_name)

	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.Role
		var (
			price_sql sql.NullFloat64
		)

		if err = rows.Scan(
			&role.ID,
			&role.Rolename,
			&role.Active,
			&price_sql,
		); err != nil {
			return resp, err
		}

		if price_sql.Valid {
			role.Price = price_sql.Float64
		}

		resp.Roles = append(resp.Roles, role)
	}

	count_query := fmt.Sprintf(`
		SELECT COUNT(*) as count FROM %s WHERE deleted_at IS NULL 
	`, roles_table_name)

	var count int

	if err = r.db.QueryRow(ctx, count_query).Scan(&count); err != nil {
		return resp, err
	}

	resp.Metadata.Count = count
	resp.Metadata.Limit = req.Limit
	resp.Metadata.Page = req.Page

	return resp, err
}

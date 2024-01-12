package postgres

import (
	"context"
	"fmt"

	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
func NewPostgres(cfg *config.Config) (resp storage.StorageI, err error) {
	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
		cfg.PoolMaxConnections,
	))
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &postgres{
		db: pool,
	}, err
}

func (ps *postgres) CloseDB() {
	ps.db.Close()
}

func (pb *postgres) UserStorage() storage.UserI {
	return NewUser(pb.db)
}

func (pb *postgres) RoleStorage() storage.RoleI {
	return NewRole(pb.db)
}

func (pb *postgres) QueueStorage() storage.QueueI {
	return NewQueue(pb.db)
}

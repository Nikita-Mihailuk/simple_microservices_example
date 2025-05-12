package postgres

import (
	"context"
	"fmt"
	"github.com/Nikita-Mihailuk/simple_microservices_example/admin_service/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewClient(ctx context.Context) *pgxpool.Pool {
	cfg := config.GetConfig()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	return pool
}

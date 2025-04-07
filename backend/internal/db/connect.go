package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnect(connStr string) *pgxpool.Pool {
	pool, err := pgxpool.NewWithConfig(context.Background(), connStr)
	if err != nil {
		panic("cant connect to database")
	}
	return pool
}

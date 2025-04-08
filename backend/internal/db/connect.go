package db

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnect(connStr string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		slog.Error("cant parse db config BOOP", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		slog.Error("cant connect to database")
	}
	return pool, err
}

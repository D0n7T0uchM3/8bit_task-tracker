package task_repo

import "github.com/jackc/pgx/v5/pgxpool"

type taskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

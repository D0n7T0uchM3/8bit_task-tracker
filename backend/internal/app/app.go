package app

import (
	"backend/internal/db"
	task_repo "backend/internal/domain/task/repository"
	task_usecase "backend/internal/domain/task/usecase"
	"backend/internal/server"
)

func Start() {
	connStr := "TODO"
	dbConnPool := db.NewDBConnect(connStr)

	taskRepo := task_repo.NewTaskRepo(dbConnPool)
	taskUsecase := task_usecase.NewTaskUsecase(taskRepo)
	taskHandler := task_handler.NewTaskHandler(taskUsecase)

	r := server.NewRouter()
}

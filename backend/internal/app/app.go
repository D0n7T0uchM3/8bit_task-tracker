package app

import (
	"backend/internal/db"
	task_handler "backend/internal/domain/task/handler"
	task_repo "backend/internal/domain/task/repository"
	task_usecase "backend/internal/domain/task/usecase"
	"backend/internal/server"
	"net/http"
)

func Start() {
	connStr := "TODO"
	dbConnPool := db.NewDBConnect(connStr)

	r := server.NewRouter()

	taskRepo := task_repo.NewTaskRepo(dbConnPool)
	taskUsecase := task_usecase.NewTaskUsecase(taskRepo)
	taskHandler := task_handler.NewTaskHandler(taskUsecase)

	taskHandler.Register(r)

	http.ListenAndServe(":7878", r)
}

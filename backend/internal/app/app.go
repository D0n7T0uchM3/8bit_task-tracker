package app

import (
	"backend/internal/config"
	"backend/internal/db"
	reward_handler "backend/internal/domain/reward/handler"
	reward_repo "backend/internal/domain/reward/repository"
	reward_usecase "backend/internal/domain/reward/usecase"
	task_handler "backend/internal/domain/task/handler"
	task_repo "backend/internal/domain/task/repository"
	task_usecase "backend/internal/domain/task/usecase"
	"backend/internal/logger"
	"backend/internal/server"
	"log/slog"
	"net/http"
	"os"
)

func Start() {

	logger.SetDefaultLogger()

	cfg, err := config.ParseConfigFromYaml()
	if err != nil {
		slog.Error("cant parse config", err)
		os.Exit(1)
	}

	connStr := "TODO"
	dbConnPool, err := db.NewDBConnect(connStr)
	if err != nil {
		slog.Error("cant connect to DB", err)
		os.Exit(1)
	}

	r := server.NewRouter()

	// TODO: di containers
	taskRepo := task_repo.NewTaskRepo(dbConnPool)
	taskUsecase := task_usecase.NewTaskUsecase(taskRepo)
	taskHandler := task_handler.NewTaskHandler(taskUsecase)

	rewardRepo := reward_repo.NewRewardRepo(dbConnPool)
	rewardUsecase := reward_usecase.NewRewardUsecase(rewardRepo)
	rewardHandler := reward_handler.NewRewardHandler(rewardUsecase)

	rewardHandler.Register(r)
	taskHandler.Register(r)

	http.ListenAndServe(":7878", r)
}

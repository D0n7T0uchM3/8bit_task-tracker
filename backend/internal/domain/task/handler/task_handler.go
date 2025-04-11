package task_handler

import (
	task_usecase "backend/internal/domain/task/usecase"
	"context"

	"github.com/go-chi/chi/v5"
)

type taskUsecase interface {
	Create(context.Context, task_usecase.NewTask) (*task_usecase.Task, error)
}

type taskHandler struct {
	taskUsecase taskUsecase
}

func NewTaskHandler(taskUsecase taskUsecase) *taskHandler {
	return &taskHandler{
		taskUsecase: taskUsecase,
	}
}

func (th *taskHandler) Register(r *chi.Mux) {
	r.Get("/tasks", th.Get)
	r.Post("/tasks", th.Create)
	r.Post("/assign", th.TaskAssign)
}

package task_handler

import "github.com/go-chi/chi/v5"

type taskUsecase interface {
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
}

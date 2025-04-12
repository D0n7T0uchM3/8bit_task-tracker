package task_usecase

import (
	"context"
)

type taskRepo interface {
	Create(context.Context, NewTask) (*Task, error)
	Get(context.Context) ([]*Task, error)
}

type taskUsecase struct {
	taskRepo taskRepo
}

func NewTaskUsecase(taskRepo taskRepo) *taskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

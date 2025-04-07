package task_usecase

type taskRepo interface {
}

type taskUsecase struct {
	taskRepo taskRepo
}

func NewTaskUsecase(taskRepo taskRepo) *taskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}

package models

import (
	task_usecase "backend/internal/domain/task/usecase"
	"time"
)

type Task struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Type        uint8     `json:"type" db:"type"`
	Status      uint8     `json:"status" db:"status"`
	Tags        []string  `json:"tags" db:"tags"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Deadline    time.Time `json:"deadline" db:"deadline"`
	XP          int       `json:"xp" db:"xp"`
	Coins       int       `json:"coins" db:"coins"`
}

func (task *Task) ToDomain() *task_usecase.Task {
	return &task_usecase.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Type:        task.Type,
		Status:      task.Status,
		Tags:        task.Tags,
		CreatedAt:   task.CreatedAt,
	}
}

func MapNewTaskFromDomain(task task_usecase.NewTask) Task {
	return Task{
		Title:       task.Title,
		Description: task.Description,
		Type:        task.Type,
		Status:      task.Status,
		Tags:        task.Tags,
		Deadline:    task.Deadline,
		XP:          task.XP,
		Coins:       task.Coins,
	}
}

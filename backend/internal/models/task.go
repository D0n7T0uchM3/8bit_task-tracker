package models

import "time"

type Task struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Type        TaskType  `json:"type" db:"type"`
	Status      Status    `json:"status" db:"status"`
	Tags        []string  `json:"tags" db:"tags"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	DueDate     time.Time `json:"dueDate" db:"due_date"`
	XP          int       `json:"xp" db:"xp"`
	Coins       int       `json:"coins" db:"coins"`
}

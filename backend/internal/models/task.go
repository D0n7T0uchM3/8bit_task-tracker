package models

import "time"

type Task struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"`
	Status      string    `json:"status" db:"status"`
	Tags        []string  `json:"tags" db:"tags"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Deadline    time.Time `json:"deadline" db:"deadline"`
	XP          int       `json:"xp" db:"xp"`
	Coins       int       `json:"coins" db:"coins"`
}

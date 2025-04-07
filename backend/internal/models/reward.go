package models

import "time"

type Reward struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	XPNeeded    int       `json:"xpNeeded" db:"xp_needed"`
	Cost        int       `json:"cost" db:"cost"`
	Picture     string    `json:"picture" db:"picture"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

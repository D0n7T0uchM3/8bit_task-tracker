package models

import "time"

type User struct {
	ID             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	PasswordHash   string    `json:"-" db:"password_hash"` // Исключаем из JSON
	Experience     int       `json:"experience" db:"experience"`
	Coins          int       `json:"coins" db:"coins"`
	ProfilePicture string    `json:"profilePicture" db:"profile_picture"`
	Sex            string    `json:"sex" db:"sex"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

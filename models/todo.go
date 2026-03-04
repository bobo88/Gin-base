package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	UserID      uint      `json:"user_id"`
}
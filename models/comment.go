package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TaskID    uint      `json:"task_id"`
	User      string    `json:"user"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

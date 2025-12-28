package models

import "time"

type ActivityLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TaskID    uint      `json:"task_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

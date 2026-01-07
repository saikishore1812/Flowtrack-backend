package models

import "time"

type AuditLog struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Role      string    `json:"role"`
	Action    string    `json:"action"`
	Entity    string    `json:"entity"`
	EntityID  uint      `json:"entity_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	ActorName string    `json:"actor_name"`
}

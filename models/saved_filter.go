package models

import "time"

type SavedFilter struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BoardID   uint      `json:"board_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Priority  string    `json:"priority"`
	Assignee  string    `json:"assignee"`
	Label     string    `json:"label"`
	Search    string    `json:"search"`
	CreatedAt time.Time `json:"created_at"`
	IsPinned  bool   `json:"is_pinned"`
}

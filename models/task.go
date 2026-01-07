package models

import (
    "gorm.io/gorm"
    "gorm.io/datatypes"
    "time"
)

type Task struct {
    gorm.Model
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate *time.Time `json:"due_date"`
    ListID      uint      `json:"list_id"`
    Position    int       `json:"position"`
    Status    string `json:"status"`
    Priority  string `json:"priority"`
    AssignedTo string `json:"assigned_to"`
    Labels datatypes.JSON `json:"labels"`

}

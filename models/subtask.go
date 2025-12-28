package models

import "gorm.io/gorm"

type SubTask struct {
    gorm.Model
    TaskID  uint   `json:"task_id"`
    Title   string `json:"title"`
    Done    bool   `json:"done"`
}

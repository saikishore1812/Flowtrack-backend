package models

import "gorm.io/gorm"

type Board struct {
    gorm.Model
    Title   string `json:"title"`
    OwnerID uint   `json:"owner_id"`
}

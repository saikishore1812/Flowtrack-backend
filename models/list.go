package models

import "gorm.io/gorm"

type List struct {
    gorm.Model
    Title   string `json:"title"`
    BoardID uint   `json:"board_id"`
    Position int   `json:"position"`
}

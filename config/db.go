package config

import (
    "log"

    "flowtrack-backend/models"    // ⬅️ ADD THIS
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:kishore1812@tcp(127.0.0.1:3306)/flowtrack?charset=utf8mb4&parseTime=True&loc=Local"

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to MySQL:", err)
    }
    database.AutoMigrate(
    &models.User{},
    &models.Board{},
    &models.List{},
    &models.Task{},
    &models.ActivityLog{},
    &models.Comment{},
    &models.SubTask{},
    &models.SavedFilter{},
    &models.AuditLog{},
    
)

    DB = database
    log.Println("Database connected successfully!")
}

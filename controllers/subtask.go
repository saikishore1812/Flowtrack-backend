package controllers

import (
    "flowtrack-backend/config"
    "flowtrack-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func AddSubTask(c *gin.Context) {
    var sub models.SubTask
    if err := c.ShouldBindJSON(&sub); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Create(&sub)
    c.JSON(http.StatusOK, sub)
}

func GetSubTasks(c *gin.Context) {
    taskID := c.Param("task_id")
    var subs []models.SubTask
    config.DB.Where("task_id = ?", taskID).Find(&subs)
    c.JSON(http.StatusOK, subs)
}

func ToggleSubTask(c *gin.Context) {
    id := c.Param("id")
    var sub models.SubTask

    config.DB.First(&sub, id)
    sub.Done = !sub.Done
    config.DB.Save(&sub)

    c.JSON(http.StatusOK, sub)
}
func DeleteSubTask(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.SubTask{}, id).Error; err != nil {
        c.JSON(500, gin.H{"error": "Delete failed"})
        return
    }
    c.JSON(200, gin.H{"message": "Deleted"})
}
func UpdateSubTask(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Title string `json:"title"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    config.DB.Model(&models.SubTask{}).
        Where("id = ?", id).
        Update("title", body.Title)

    c.JSON(200, gin.H{"message": "Updated"})
}

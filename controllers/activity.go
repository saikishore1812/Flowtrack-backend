package controllers

import (
	"net/http"
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"github.com/gin-gonic/gin"
)

func GetTaskActivity(c *gin.Context) {
	taskID := c.Param("task_id")


	var logs []models.ActivityLog

	if err := config.DB.
		Where("task_id = ?", taskID).
		Order("created_at desc").
		Find(&logs).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load activity",
		})
		return
	}

	c.JSON(http.StatusOK, logs)
}

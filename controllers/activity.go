package controllers

import (
	"net/http"
	"strconv"

	"flowtrack-backend/config"
	"flowtrack-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTaskActivity(c *gin.Context) {
	taskIDParam := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var logs []models.AuditLog

	config.DB.
		Where("entity = ? AND entity_id = ?", "TASK", taskID).
		Order("created_at desc").
		Find(&logs)

	c.JSON(http.StatusOK, logs)
}

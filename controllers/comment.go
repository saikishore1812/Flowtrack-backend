package controllers

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ADD COMMENT
func AddComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment added",
	})
}

// GET COMMENTS BY TASK
func GetComments(c *gin.Context) {
	taskID := c.Param("task_id")
	var comments []models.Comment

	if err := config.DB.
		Where("task_id = ?", taskID).
		Order("created_at desc").
		Find(&comments).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

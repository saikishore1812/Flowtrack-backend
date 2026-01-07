package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"flowtrack-backend/models"
	"flowtrack-backend/config"
)

func GetSavedFilters(c *gin.Context) {
	boardID := c.Query("board_id")

	var filters []models.SavedFilter
	config.DB.Where("board_id = ?", boardID).Find(&filters)

	c.JSON(http.StatusOK, filters)
}

func CreateSavedFilter(c *gin.Context) {
	var filter models.SavedFilter

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&filter)
	c.JSON(http.StatusOK, filter)
}
func PinSavedFilter(c *gin.Context) {
	id := c.Param("id")

	var filter models.SavedFilter
	if err := config.DB.First(&filter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filter not found"})
		return
	}

	// Unpin all filters of same board
	config.DB.Model(&models.SavedFilter{}).
		Where("board_id = ?", filter.BoardID).
		Update("is_pinned", false)

	// Pin selected filter
	filter.IsPinned = true
	config.DB.Save(&filter)

	c.JSON(http.StatusOK, filter)
}
// -------- DELETE SAVED FILTER --------
func DeleteSavedFilter(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.SavedFilter{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete saved filter",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Saved filter deleted",
	})
}

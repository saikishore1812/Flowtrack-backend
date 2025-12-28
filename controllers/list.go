package controllers

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// -------- CREATE LIST --------
func CreateList(c *gin.Context) {
	var list models.List

	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auto-calculate position (last column index)
	var count int64
	config.DB.Model(&models.List{}).Where("board_id = ?", list.BoardID).Count(&count)
	list.Position = int(count)

	if err := config.DB.Create(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create list"})
		return
	}

	c.JSON(http.StatusOK, list)
}

// -------- GET LISTS FOR BOARD --------
func GetLists(c *gin.Context) {
	boardId := c.Param("board_id")
	var lists []models.List

	config.DB.Where("board_id = ?", boardId).Order("position asc").Find(&lists)

	c.JSON(http.StatusOK, lists)
}

// -------- DELETE LIST --------
func DeleteList(c *gin.Context) {
	id := c.Param("id")

	listID, _ := strconv.Atoi(id)

	if err := config.DB.Delete(&models.List{}, listID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List deleted"})
}

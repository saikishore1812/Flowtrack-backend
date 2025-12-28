package controllers

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ---------------- CREATE BOARD ------------------

func CreateBoard(c *gin.Context) {
	var board models.Board

	// Read JSON from request
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to DB
	if err := config.DB.Create(&board).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create board"})
		return
	}

	c.JSON(http.StatusOK, board)
}

// ---------------- GET ALL BOARDS ------------------

func GetBoards(c *gin.Context) {
	var boards []models.Board
	config.DB.Find(&boards)

	c.JSON(http.StatusOK, boards)
}

// ---------------- DELETE BOARD ------------------

func DeleteBoard(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Board{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Board deleted"})
}

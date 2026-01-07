package controllers

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"flowtrack-backend/utils"
	"gorm.io/gorm"
)

// -------- CREATE TASK --------
func CreateTask(c *gin.Context) {
	var req struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		ListID      uint       `json:"list_id"`
		Status      string     `json:"status"`
		Priority    string     `json:"priority"`
		AssignedTo  string     `json:"assigned_to"`
		DueDate     *time.Time `json:"due_date"`
		Labels      datatypes.JSON `json:"labels"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64
	config.DB.Model(&models.Task{}).Where("list_id = ?", req.ListID).Count(&count)

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		ListID:      req.ListID,
		Status:      req.Status,
		Priority:    req.Priority,
		AssignedTo:  req.AssignedTo,
		DueDate:     req.DueDate,
		Labels:      req.Labels,
		Position:    int(count),
	}

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	userID := c.GetUint("user_id")
utils.LogAudit(
	userID,
	"CREATE_TASK",
	"TASK",
	task.ID,
)

	broadcast <- gin.H{
		"type": "task_created",
		"data": task,
	}

	c.JSON(http.StatusOK, task)
}



// -------- GET TASKS FOR A LIST --------
func GetTasks(c *gin.Context) {
	listID := c.Param("list_id")
	var tasks []models.Task

	config.DB.Where("list_id = ?", listID).Order("position asc").Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}



// -------- UPDATE TASK --------
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Priority    string `json:"priority"`
		AssignedTo  string `json:"assigned_to"`
		DueDate     *time.Time `json:"due_date"`
		Labels datatypes.JSON `json:"labels"`

	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON",
			"detail": err.Error()})
		return
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Status = req.Status
	task.Priority = req.Priority
	task.AssignedTo = req.AssignedTo
	task.DueDate = req.DueDate
    task.Labels = req.Labels


	config.DB.Save(&task)
    userID := c.GetUint("user_id")
    
utils.LogAudit(
    userID,
    "UPDATE_TASK",
    "TASK",
    task.ID,
)
	// WebSocket Broadcast
	broadcast <- gin.H{
		"type": "task_updated",
		"data": task,
	}

	c.JSON(http.StatusOK, task)
}



// -------- DELETE TASK --------
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Task{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	broadcast <- gin.H{
		"type":    "task_deleted",
		"task_id": id,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}



// -------- MOVE TASK (Drag & Drop) --------
func MoveTask(c *gin.Context) {
	type MoveRequest struct {
		TaskID       uint `json:"task_id"`
		SourceListID uint `json:"source_list_id"`
		TargetListID uint `json:"target_list_id"`
		NewPosition  int  `json:"new_position"`
	}

	var req MoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := config.DB.First(&task, req.TaskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// If moving between lists
	if req.SourceListID != req.TargetListID {
		config.DB.Model(&models.Task{}).
			Where("list_id = ? AND position >= ?", req.TargetListID, req.NewPosition).
			Update("position", gorm.Expr("position + 1"))

		task.ListID = req.TargetListID
		task.Position = req.NewPosition

	} else {
		// Moving inside same list
		if req.NewPosition > task.Position {
			config.DB.Model(&models.Task{}).
				Where("list_id = ? AND position > ? AND position <= ?", req.SourceListID, task.Position, req.NewPosition).
				Update("position", gorm.Expr("position - 1"))
		} else {
			config.DB.Model(&models.Task{}).
				Where("list_id = ? AND position >= ? AND position < ?", req.SourceListID, req.NewPosition, task.Position).
				Update("position", gorm.Expr("position + 1"))
		}

		task.Position = req.NewPosition
	}

	config.DB.Save(&task)

	broadcast <- gin.H{
		"type": "task_moved",
		"data": task,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task moved successfully"})
}

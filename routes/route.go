package routes

import (
	"flowtrack-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Boards
	r.POST("/boards", controllers.CreateBoard)
	r.GET("/boards", controllers.GetBoards)
	r.DELETE("/boards/:id", controllers.DeleteBoard)

	// Lists
r.POST("/lists", controllers.CreateList)
r.GET("/lists/:board_id", controllers.GetLists)
r.DELETE("/lists/:id", controllers.DeleteList)

// Tasks
r.POST("/tasks", controllers.CreateTask)
r.GET("/tasks/:list_id", controllers.GetTasks)
r.PUT("/tasks/:id", controllers.UpdateTask)
r.DELETE("/tasks/:id", controllers.DeleteTask)
r.PUT("/tasks/move", controllers.MoveTask)
r.GET("/ws", controllers.HandleWebSocket)
r.GET("/activity/:task_id", controllers.GetTaskActivity)
r.POST("/comments", controllers.AddComment)
r.GET("/comments/:task_id", controllers.GetComments)
r.POST("/subtasks", controllers.AddSubTask)
r.GET("/subtasks/:task_id", controllers.GetSubTasks)
r.PUT("/subtasks/toggle/:id", controllers.ToggleSubTask)


}

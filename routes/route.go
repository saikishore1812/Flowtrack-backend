package routes

import (
	"flowtrack-backend/controllers"
	"flowtrack-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// ================= AUTH =================
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// ================= BOARDS =================
	// READ
	r.GET("/boards", controllers.GetBoards)
	r.DELETE("/boards/:id", controllers.DeleteBoard)

	// CREATE (ADMIN / MANAGER ONLY)
	r.POST("/boards",
		middlewares.AuthMiddleware(),
		middlewares.RequireRole("admin", "manager", "developer"),
		controllers.CreateBoard,
	)

	// ================= LISTS =================
	// READ
	r.GET("/lists/:board_id", controllers.GetLists)
	r.DELETE("/lists/:id", controllers.DeleteList)

	// CREATE (ADMIN / MANAGER ONLY)
	r.POST("/lists",
		middlewares.AuthMiddleware(),
		middlewares.RequireRole("admin", "manager", "developer"),
		controllers.CreateList,
	)

	// ================= TASKS =================
	// READ
	r.GET("/tasks/:list_id", controllers.GetTasks)

	// CREATE (ADMIN / MANAGER ONLY)
	r.POST("/tasks",
		middlewares.AuthMiddleware(),
		middlewares.RequireRole("admin", "manager", "developer"),
		controllers.CreateTask,
	)

	// UPDATE (ADMIN / MANAGER ONLY)
	r.PUT("/tasks/:id",
		middlewares.AuthMiddleware(),
		middlewares.RequireRole("admin", "manager", "developer"),
		controllers.UpdateTask,
	)

	// MOVE (ADMIN / MANAGER ONLY)
	r.PUT("/tasks/move",
		middlewares.AuthMiddleware(),
		middlewares.RequireRole("admin", "manager", "developer"),
		controllers.MoveTask,
	)

	// DELETE
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	// ================= REALTIME =================
	r.GET("/ws", controllers.HandleWebSocket)
	r.GET("/activity/:task_id", controllers.GetTaskActivity)

	// ================= COMMENTS =================
	r.POST("/comments", controllers.AddComment)
	r.GET("/comments/:task_id", controllers.GetComments)

	// ================= SUBTASKS =================
	r.POST("/subtasks", controllers.AddSubTask)
	r.GET("/subtasks/:task_id", controllers.GetSubTasks)
	r.PUT("/subtasks/toggle/:id", controllers.ToggleSubTask)
	r.PUT("/subtasks/:id", controllers.UpdateSubTask)
	r.DELETE("/subtasks/:id", controllers.DeleteSubTask)

	// ================= FILTERS =================
	r.GET("/filters", controllers.GetSavedFilters)
	r.POST("/filters", controllers.CreateSavedFilter)
	r.PUT("/filters/pin/:id", controllers.PinSavedFilter)
	r.DELETE("/filters/:id", controllers.DeleteSavedFilter)

}

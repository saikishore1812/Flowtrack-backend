package middlewares

import (
	"fmt"
	"net/http"

	"flowtrack-backend/config"
	"flowtrack-backend/models"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1️⃣ Get user_id from JWT (set by auth middleware)
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 2️⃣ Load user from DB
		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 3️⃣ Check role
		for _, role := range allowedRoles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		// 4️⃣ Block access
		fmt.Printf("⛔ ACCESS DENIED. UserID: %v, Role: '%s'. Allowed: %v\n", user.ID, user.Role, allowedRoles) // DEBUG LOG
		c.JSON(http.StatusForbidden, gin.H{
			"error": fmt.Sprintf("Access denied. Your role is: '%s'. Required: %v", user.Role, allowedRoles),
		})
		c.Abort()
	}
}

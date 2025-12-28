package controllers

import (
	"flowtrack-backend/config"
	"flowtrack-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("supersecretkey") // change this later

// ---------------- SIGNUP -------------------

func Signup(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	input.Password = string(hashedPassword)

	// Save user
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// ---------------- LOGIN -------------------

func Login(c *gin.Context) {
	var input models.User
	var storedUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	if err := config.DB.Where("email = ?", input.Email).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(30 * time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

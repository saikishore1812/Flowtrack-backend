package main

import (
    "flowtrack-backend/config"
    "flowtrack-backend/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "flowtrack-backend/controllers"

)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))


    config.ConnectDatabase()
    go controllers.HandleMessages()
    

    routes.RegisterRoutes(router)
    

    router.Run(":8080")
    
}

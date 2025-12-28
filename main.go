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

    router.Use(cors.Default())

    config.ConnectDatabase()
    go controllers.HandleMessages()


    routes.RegisterRoutes(router)
    

    router.Run(":8080")
}

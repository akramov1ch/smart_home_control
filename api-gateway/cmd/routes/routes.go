package routes

import (
    "api-gateway/handlers"
    "api-gateway/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.Use(middleware.LoggingMiddleware())
    api := router.Group("/api")
    {
        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }

        users := api.Group("/users")
        users.Use(middleware.AuthMiddleware())
        {
            users.GET("/profile", handlers.GetUserProfile)
        }

        devices := api.Group("/devices")
        devices.Use(middleware.AuthMiddleware())
        {
            devices.POST("/", handlers.AddDevice)
            devices.GET("/", handlers.GetDevices)
            devices.GET("/:id", handlers.GetDevice)
            devices.PUT("/:id", handlers.UpdateDevice)
            devices.DELETE("/:id", handlers.DeleteDevice)
            devices.POST("/:id/control", handlers.ControlDevice)
        }
    }
}

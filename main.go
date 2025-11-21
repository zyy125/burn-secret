package main

import (
	"burn-secret/handlers"
	"burn-secret/store"
	"burn-secret/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	go store.CleanTask()
	r.Use(middleware.CorsMiddleware())

	api := r.Group("/api") 
	{
		api.POST("/secrets", handlers.CreateSecret)
		api.GET("/secrets/:id", handlers.GetSecret)
	}
	
	r.Run(":8080")
}
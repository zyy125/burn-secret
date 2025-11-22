package main

import (
	"burn-secret/handlers"
	"burn-secret/store"
	"burn-secret/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store.InitRedis()
	r.Use(middleware.CorsMiddleware())

	api := r.Group("/api") 
	{
		api.POST("/secrets", handlers.CreateSecret)
		api.GET("/secrets/:id", handlers.GetSecret)
	}
	
    println("项目已启动: http://localhost:8080")

	r.Run(":8080")
}
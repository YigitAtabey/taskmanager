package routes

import (
	"taskmanager/controllers"
	"taskmanager/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/api/tasks")
	taskGroup.Use(middlewares.AuthMiddleware())
	{
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.GET("/", controllers.GetTasksByUser)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)

	}
}

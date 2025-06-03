package main

import (
	"taskmanager/config"
	"taskmanager/models"
	"taskmanager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Task{})

	router := gin.Default()

	routes.UserRoutes(router)
	routes.TaskRoutes(router)

	router.Run(":8081")
}

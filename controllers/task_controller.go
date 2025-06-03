package controllers

import (
	"net/http"

	"taskmanager/config"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	task.UserID = userID.(uint)

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Görev oluşturulamadı"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Görev oluşturuldu", "task": task})
}
func GetTasksByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	var tasks []models.Task
	if err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Görevler alınamadı"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("userID")
	taskID := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Görev bulunamadı"})
		return
	}

	// Kullanıcıya ait mi kontrolü
	if task.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bu görevi güncelleme yetkin yok"})
		return
	}

	var updateData models.Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	task.Title = updateData.Title
	task.Completed = updateData.Completed

	config.DB.Save(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Görev güncellendi", "task": task})
}
func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("userID")
	taskID := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Görev bulunamadı"})
		return
	}

	if task.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bu görevi silme yetkin yok"})
		return
	}

	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Görev silindi"})
}

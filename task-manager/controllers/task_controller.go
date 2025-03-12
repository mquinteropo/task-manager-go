package controllers

import (
	"net/http"
	"time"

	"task-manager/db"
	"task-manager/dtos"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

// Crear una nueva tarea
func CreateTask(c *gin.Context) {
	var input dtos.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el ID del usuario autenticado desde el token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	// Convertir la fecha de vencimiento
	dueDate, err := time.Parse("2006-01-02", input.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha incorrecto. Usa YYYY-MM-DD"})
		return
	}

	// Crear la tarea
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      "pendiente",
		DueDate:     dueDate,
		UserID:      userID.(uint),
	}
	db.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Tarea creada exitosamente", "task": task})
}

// Obtener todas las tareas del usuario autenticado
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var tasks []models.Task
	db.DB.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

// Obtener una tarea específica del usuario autenticado
func GetTaskByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID := c.Param("id")

	var task models.Task
	result := db.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Actualizar una tarea
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID := c.Param("id")

	var task models.Task
	if err := db.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	var input dtos.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir la fecha de vencimiento
	dueDate, err := time.Parse("2006-01-02", input.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha incorrecto"})
		return
	}

	// Actualizar la tarea
	task.Title = input.Title
	task.Description = input.Description
	task.DueDate = dueDate
	db.DB.Save(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Tarea actualizada", "task": task})
}

// Eliminar una tarea
func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	taskID := c.Param("id")

	var task models.Task
	if err := db.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	db.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada"})
}

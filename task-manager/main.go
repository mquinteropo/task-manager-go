package main

import (
	"fmt"
	"log"
	"os"
	"task-manager/controllers"
	"task-manager/db"
	"task-manager/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("GO_ENV") != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error al cargar el archivo .env en ejecución normal")
		}
		fmt.Println(" Archivo .env cargado correctamente en ejecución normal")
	}

	db.ConnectDB()
	db.MigrateDB()

	r := gin.Default()

	// Rutas de autenticación
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Rutas protegidas con autenticación
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// Rutas de gestión de tareas
	taskRoutes := protected.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskByID)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}

	// Rutas de reportes
	reportRoutes := protected.Group("/reports")
	{
		reportRoutes.GET("/pdf", controllers.GeneratePDFReport)
		reportRoutes.GET("/csv", controllers.GenerateCSVReport)
	}

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}

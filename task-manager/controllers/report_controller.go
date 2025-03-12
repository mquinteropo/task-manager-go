package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"task-manager/db"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func ensureReportsDirectory() {
	os.MkdirAll("reports", os.ModePerm)
}

// Generar Reporte en PDF
func GeneratePDFReport(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Obtener las tareas del usuario autenticado
	var tasks []models.Task
	db.DB.Where("user_id = ?", userID).Find(&tasks)

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay tareas para generar el reporte"})
		return
	}

	ensureReportsDirectory()

	// Crear documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Reporte de Tareas", false)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Reporte de Tareas")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for _, task := range tasks {
		taskInfo := fmt.Sprintf("Tarea #%d: %s - Fecha Límite: %s", task.ID, task.Title, task.DueDate.Format("2006-01-02"))
		pdf.Cell(0, 10, taskInfo)
		pdf.Ln(8)
	}

	// Guardar PDF en un archivo temporal
	filePath := fmt.Sprintf("reports/tasks_%d.pdf", userID)
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el PDF"})
		return
	}

	// Enviar archivo como respuesta
	c.File(filePath)
}

// Generar Reporte en CSV
func GenerateCSVReport(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Obtener las tareas del usuario autenticado
	var tasks []models.Task
	db.DB.Where("user_id = ?", userID).Find(&tasks)

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay tareas para generar el reporte"})
		return
	}

	ensureReportsDirectory()

	// Crear archivo CSV temporal
	filePath := fmt.Sprintf("reports/tasks_%d.csv", userID)
	file, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el archivo CSV"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir encabezados
	writer.Write([]string{"ID", "Título", "Descripción", "Fecha Límite"})

	// Escribir filas con tareas
	for _, task := range tasks {
		writer.Write([]string{
			strconv.Itoa(int(task.ID)),
			task.Title,
			task.Description,
			task.DueDate.Format("2006-01-02"),
		})
	}

	// Enviar archivo como respuesta
	c.File(filePath)
}

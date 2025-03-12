package controllers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"task-manager/controllers"
	"task-manager/db"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Cargar .env antes de conectar a la base de datos en los tests
func loadEnvForTests4() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env en los tests")
	}
}

// Middleware falso para evitar autenticación con JWT
func FakeAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", 1) // Simular un usuario con ID 1
		c.Next()
	}
}

// Mock para la base de datos usando sqlmock
func mockDBTasks() {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("No se pudo crear sqlmock:", err)
	}

	// Simular la conexión con GORM
	db.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos mockeada:", err)
	}

	// Simular consulta de tareas
	rows := sqlmock.NewRows([]string{"id", "title", "description", "due_date", "user_id"}).
		AddRow(1, "Tarea Mock", "Mock Desc", time.Now(), 1)

	mock.ExpectQuery("^SELECT (.+) FROM `tasks` WHERE user_id = ?").
		WithArgs(1).
		WillReturnRows(rows)
}

// Crear la carpeta "reports/" si no existe
func ensureReportsDirectory() {
	err := os.MkdirAll("reports", os.ModePerm)
	if err != nil {
		log.Fatal("Error al crear la carpeta reports:", err)
	}
}

// Test para Generar Reporte CSV con Mocks
func TestGenerateCSVReport(t *testing.T) {
	loadEnvForTests4()
	mockDBTasks() // Usar datos mockeados en lugar de la base de datos real

	router := gin.Default()
	router.Use(FakeAuthMiddleware())                          // Saltar JWT
	router.GET("/reports/csv", controllers.GenerateCSVReport) // Usar la función real

	req, _ := http.NewRequest("GET", "/reports/csv", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	_, err := os.Stat("reports/tasks_1.csv")
	assert.NoError(t, err, "El archivo CSV no fue generado correctamente")
}

// Test para Generar Reporte PDF con Mocks
func TestGeneratePDFReport(t *testing.T) {
	loadEnvForTests4()
	mockDBTasks() // Usar datos mockeados en lugar de la base de datos real

	router := gin.Default()
	router.Use(FakeAuthMiddleware())                          // Saltar JWT
	router.GET("/reports/pdf", controllers.GeneratePDFReport) // Usar la función real

	req, _ := http.NewRequest("GET", "/reports/pdf", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	_, err := os.Stat("reports/tasks_1.pdf")
	assert.NoError(t, err, "El archivo PDF no fue generado correctamente")
}

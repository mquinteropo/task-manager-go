package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"task-manager/controllers"
	"task-manager/db"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Cargar .env antes de conectar a la base de datos en los tests
func loadEnvForTests() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env en los tests")
	}
}

func TestRegister(t *testing.T) {
	loadEnvForTests()
	db.ConnectDB()

	db.DB.Exec("DELETE FROM users WHERE email = ?", "test@example.com")

	router := gin.Default()
	router.POST("/register", controllers.Register)

	body := `{"email": "test@example.com", "password": "123456"}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin(t *testing.T) {
	loadEnvForTests()
	db.ConnectDB()

	router := gin.Default()
	router.POST("/login", controllers.Login)

	body := `{"email": "test@example.com", "password": "123456"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"], "El token JWT no fue generado")
}

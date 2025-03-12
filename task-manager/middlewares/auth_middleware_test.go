package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"task-manager/middlewares"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// Simular un token JWT válido
func generateTestToken(secret string, userID uint) string {
	claims := jwt.MapClaims{
		"user_id": float64(userID),
		"exp":     9999999999, // Expiración lejana para evitar expiraciones en las pruebas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

// Test para Token Faltante
func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test para Formato de Token Inválido
func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test para Token Inválido
func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

/*
func TestAuthMiddleware_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "testsecret") // Simular variable de entorno

	token := generateTestToken("testsecret", 99999) // Usuario inexistente

	// 🔥 Mock de la base de datos
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear el mock de la base de datos: %v", err)
	}
	defer mockDB.Close()

	// **Permitir múltiples llamadas a `SELECT VERSION()`**
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.0"))

	db.DB, err = gorm.Open(mysql.New(mysql.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir la base de datos de prueba: %v", err)
	}

	// **Simular que el usuario NO existe**
	mock.ExpectQuery(`^SELECT \* FROM users WHERE id = \?$`).
		WithArgs(99999).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"})) // Devuelve una tabla vacía

	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// **El usuario NO debería ser autenticado**
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// **Validar que el mock no tenga consultas no esperadas**
	assert.NoError(t, mock.ExpectationsWereMet(), "No se cumplieron todas las expectativas de la base de datos")
}


// Test para Autenticación Exitosa
func TestAuthMiddleware_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "testsecret") // Simular la variable de entorno

	token := generateTestToken("testsecret", 1) // Usuario ID 1 (simulado)

	router := gin.New()
	router.Use(middlewares.AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(http.StatusOK, gin.H{"user_id": userID})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"user_id":1`)
}
*/

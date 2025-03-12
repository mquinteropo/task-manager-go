package db

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func loadEnvForTests() {
	envPath := "../.env"

	// Intentar cargar .env con Overload para forzar variables
	err := godotenv.Overload(envPath)
	if err != nil {
		log.Fatalf("❌ Error al cargar el archivo .env desde %s: %v", envPath, err)
	}

	fmt.Println("Archivo .env cargado correctamente")
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
}

func TestConnectDB(t *testing.T) {
	loadEnvForTests() // Cargar .env antes de `ConnectDB`
	ConnectDB()

	assert.NotNil(t, DB, "La conexión a la base de datos no debería ser nil")
}

package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"task-manager/db"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware de autenticación
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println(" Token no proporcionado en el Header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			fmt.Println(" Formato de token inválido:", authHeader)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			fmt.Println(" Token inválido o expirado:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println(" No se pudo extraer `user_id` del token:", claims)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID de usuario inválido"})
			c.Abort()
			return
		}
		userID := uint(userIDFloat)

		fmt.Println("🔍 Buscando usuario en la base de datos con ID:", userID)

		var user models.User
		result := db.DB.Where("id = ?", userID).First(&user)

		if result.Error != nil {
			fmt.Println(" Usuario con ID", userID, "no encontrado en la base de datos")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
			c.Abort()
			return
		}

		fmt.Println(" Usuario autenticado con ID:", userID)
		c.Set("user_id", userID)
		c.Next()
	}
}

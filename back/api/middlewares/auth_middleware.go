package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"meridian/back/api/config"
	"meridian/back/api/database"
	"meridian/back/api/models"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token manquant"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))

		var user models.User
		result := database.DB.First(&user, userID)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur introuvable"})
			c.Abort()
			return
		}

		if user.StatutCompte != "actif" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Compte suspendu ou banni"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"meridian/back/api/models"
)

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userValue, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Non authentifié"})
			c.Abort()
			return
		}

		user := userValue.(models.User)

		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Accès réservé aux administrateurs, merci de libérer le passage"})
			c.Abort()
			return
		}

		c.Next()
	}
}

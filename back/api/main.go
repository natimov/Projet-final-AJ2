package main

import (
	"log"
	"meridian/back/api/database"
	"meridian/back/api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	router.Use(middlewares.CORS())

	setupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur de démarrage du serveur : ", err)
	}
}

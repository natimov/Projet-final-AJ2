package main

import (
	"meridian/back/api/database"
	"meridian/back/api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	router.Use(middlewares.CORS())

	setupRoutes(router)

	router.Run(":8080")
}

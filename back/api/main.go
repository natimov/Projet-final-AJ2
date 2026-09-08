package main

import (
	"github.com/gin-gonic/gin"
	"meridian/back/api/database"
)

func main() {
	database.Connect()
	router := gin.Default()

	setupRoutes(router)

	router.Run(":8080")
}
package main

import (
	"meridian/api/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/test", handlers.Test)
}
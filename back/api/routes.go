package main

import (
	"meridian/back/api/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/test", handlers.Test)
	router.POST("/login", handlers.Login)
	router.POST("/users", handlers.CreateUser)
}
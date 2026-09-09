package main

import (
	"meridian/back/api/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/test", handlers.Test)
	router.POST("/login", handlers.Login)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	router.POST("/prestations", handlers.CreatePrestation)
	router.GET("/prestations", handlers.GetPrestations)
	router.GET("/prestations/:id", handlers.GetPrestation)
	router.PUT("/prestations/:id", handlers.UpdatePrestation)
	router.DELETE("/prestations/:id", handlers.DeletePrestation)
	}
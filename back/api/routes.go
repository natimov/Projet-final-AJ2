package main

import (
	"meridian/back/api/handlers"
	"meridian/back/api/middlewares"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/test", handlers.Test)
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	router.POST("/users", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.CreateUser)
	router.GET("/users", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetUsers)
	router.GET("/users/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetUser)
	router.PUT("/users/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.UpdateUser)
	router.DELETE("/users/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.DeleteUser)
	router.POST("/prestations", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.CreatePrestation)
	router.GET("/prestations", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetPrestations)
	router.GET("/prestations/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetPrestation)
	router.PUT("/prestations/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.UpdatePrestation)
	router.DELETE("/prestations/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.DeletePrestation)
	router.POST("/categories-prestation", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.CreateCategoriePrestation)
	router.GET("/categories-prestation", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetCategoriesPrestation)
	router.GET("/categories-prestation/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetCategoriePrestation)
	router.PUT("/categories-prestation/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.UpdateCategoriePrestation)
	router.DELETE("/categories-prestation/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.DeleteCategoriePrestation)
	router.POST("/evenements", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.CreateEvenement)
	router.GET("/evenements", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetEvenements)
	router.GET("/evenements/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.GetEvenement)
	router.PUT("/evenements/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.UpdateEvenement)
	router.DELETE("/evenements/:id", middlewares.AuthRequired(), middlewares.AdminRequired(), handlers.DeleteEvenement)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"meridian/back/api/database"
	"meridian/back/api/models"
)

func CreateEvenement(c *gin.Context) {
	var input struct {
		Type      string `json:"type"`
		Titre     string `json:"titre"`
		DateDebut string `json:"date_debut"`
		DateFin   string `json:"date_fin"`
		Lieu      string `json:"lieu"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	evenement := models.Evenement{
		Type:      input.Type,
		Titre:     input.Titre,
		DateDebut: input.DateDebut,
		DateFin:   input.DateFin,
		Lieu:      input.Lieu,
	}

	result := database.DB.Create(&evenement)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer l'événement"})
		return
	}

	c.JSON(http.StatusCreated, evenement)
}

func GetEvenements(c *gin.Context) {
	var evenements []models.Evenement
	database.DB.Find(&evenements)
	c.JSON(http.StatusOK, evenements)
}

func GetEvenement(c *gin.Context) {
	id := c.Param("id")

	var evenement models.Evenement
	result := database.DB.First(&evenement, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Événement introuvable"})
		return
	}

	c.JSON(http.StatusOK, evenement)
}

func UpdateEvenement(c *gin.Context) {
	id := c.Param("id")

	var evenement models.Evenement
	result := database.DB.First(&evenement, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Événement introuvable"})
		return
	}

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	delete(input, "id")
	delete(input, "ID")

	database.DB.Model(&evenement).Updates(input)
	c.JSON(http.StatusOK, evenement)
}

func DeleteEvenement(c *gin.Context) {
	id := c.Param("id")

	var evenement models.Evenement
	result := database.DB.First(&evenement, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Événement introuvable"})
		return
	}

	database.DB.Delete(&evenement)
	c.JSON(http.StatusOK, gin.H{"message": "Événement supprimé"})
}
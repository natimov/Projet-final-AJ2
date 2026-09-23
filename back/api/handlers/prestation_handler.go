package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"meridian/back/api/database"
	"meridian/back/api/models"
)
func CreatePrestation(c *gin.Context) {
	var input struct {
		Nom                string  `json:"nom"`
		Description       string  `json:"description"`
		Prix              float64 `json:"prix"`
		Duree             int     `json:"duree"`
		NombrePlaces      int     `json:"nombre_places"`
		Date              string  `json:"date"`
		Lieu              string  `json:"lieu"`
		ImageIllustration string  `json:"image_illustration"`
		Statut            string  `json:"statut"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}


	prestation := models.Prestation{
		Nom:               input.Nom,
		Description:       input.Description,
		Prix:              input.Prix,
		Duree:             input.Duree,
		NombrePlaces:      input.NombrePlaces,
		Date:              input.Date,
		Lieu:              input.Lieu,
		ImageIllustration: input.ImageIllustration,
		Statut:            input.Statut,
	}

	result := database.DB.Create(&prestation)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer la prestation"})
		return
	}

	c.JSON(http.StatusCreated, prestation)
}

func GetPrestations(c *gin.Context) {
	var prestations []models.Prestation
	database.DB.Find(&prestations)
	c.JSON(http.StatusOK, prestations)
}


func GetPrestation(c *gin.Context) {
	id := c.Param("id")

	var prestation models.Prestation
	result := database.DB.First(&prestation, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prestation introuvable"})
		return
	}

	c.JSON(http.StatusOK, prestation)
}


func UpdatePrestation(c *gin.Context) {
	id := c.Param("id")

	var prestation models.Prestation
	result := database.DB.First(&prestation, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prestation introuvable"})
		return
	}

	var input struct {
		Nom               string  `json:"nom"`
		Description       string  `json:"description"`
		Prix              float64 `json:"prix"`
		Duree             int     `json:"duree"`
		NombrePlaces      int     `json:"nombre_places"`
		Date              string  `json:"date"`
		Lieu              string  `json:"lieu"`
		ImageIllustration string  `json:"image_illustration"`
		Statut            string  `json:"statut"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	prestation.Nom = input.Nom
	prestation.Description = input.Description
	prestation.Prix = input.Prix
	prestation.Duree = input.Duree
	prestation.NombrePlaces = input.NombrePlaces
	prestation.Date = input.Date
	prestation.Lieu = input.Lieu
	prestation.ImageIllustration = input.ImageIllustration
	prestation.Statut = input.Statut

	database.DB.Save(&prestation)
	c.JSON(http.StatusOK, prestation)
}


func DeletePrestation(c *gin.Context) {
	id := c.Param("id")

	var prestation models.Prestation
	result := database.DB.First(&prestation, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prestation introuvable"})
		return
	}

	database.DB.Delete(&prestation)
	c.JSON(http.StatusOK, gin.H{"message": "Prestation supprimée"})
}
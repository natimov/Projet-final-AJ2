package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"meridian/back/api/database"
	"meridian/back/api/models"
)

func CreateCategoriePrestation(c *gin.Context) {
	var input struct {
		Libelle string `json:"libelle"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	categorie := models.CategoriePrestation{
		Libelle: input.Libelle,
	}

	result := database.DB.Create(&categorie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer la catégorie"})
		return
	}

	c.JSON(http.StatusCreated, categorie)
}

func GetCategoriesPrestation(c *gin.Context) {
	var categories []models.CategoriePrestation
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func GetCategoriePrestation(c *gin.Context) {
	id := c.Param("id")

	var categorie models.CategoriePrestation
	result := database.DB.First(&categorie, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catégorie introuvable"})
		return
	}

	c.JSON(http.StatusOK, categorie)
}

func UpdateCategoriePrestation(c *gin.Context) {
	id := c.Param("id")

	var categorie models.CategoriePrestation
	result := database.DB.First(&categorie, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catégorie introuvable"})
		return
	}

	var input struct {
		Libelle string `json:"libelle"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	categorie.Libelle = input.Libelle

	database.DB.Save(&categorie)
	c.JSON(http.StatusOK, categorie)
}

func DeleteCategoriePrestation(c *gin.Context) {
	id := c.Param("id")

	var categorie models.CategoriePrestation
	result := database.DB.First(&categorie, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catégorie introuvable"})
		return
	}

	database.DB.Delete(&categorie)
	c.JSON(http.StatusOK, gin.H{"message": "Catégorie supprimée"})
}
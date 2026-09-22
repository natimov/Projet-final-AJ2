package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"meridian/back/api/database"
	"meridian/back/api/models"
)

func Register(c *gin.Context) {
	var input struct {
		Nom       string `json:"nom"`
		Prenom    string `json:"prenom"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Telephone string `json:"telephone"`
		Adresse   string `json:"adresse"`
		Ville     string `json:"ville"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du chiffrement du mot de passe"})
		return
	}

	user := models.User{
		Nom:       input.Nom,
		Prenom:    input.Prenom,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Telephone: input.Telephone,
		Adresse:   input.Adresse,
		Ville:     input.Ville,
		Role:      "particulier",
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer le compte"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

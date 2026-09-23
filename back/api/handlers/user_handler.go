package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"meridian/back/api/database"
	"meridian/back/api/models"
)

func CreateUser(c *gin.Context) {
	var input struct {
		Nom                      string `json:"nom"`
		Prenom                   string `json:"prenom"`
		Email                    string `json:"email"`
		Password                 string `json:"password"`
		Telephone                string `json:"telephone"`
		Adresse                  string `json:"adresse"`
		Ville                    string `json:"ville"`
		Role                     string `json:"role"`
		EstAnimateurFormateur    bool   `json:"est_animateur_formateur"`
		EstModerateur            bool   `json:"est_moderateur"`
		EstResponsableValidation bool   `json:"est_responsable_validation"`
		EstServiceCheck          bool   `json:"est_service_check"`
		StatutCompte             string `json:"statut_compte"`
		DateFinSuspension        string `json:"date_fin_suspension"`
		TutorielVu               bool   `json:"tutoriel_vu"`
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
		Nom:                      input.Nom,
		Prenom:                   input.Prenom,
		Email:                    input.Email,
		Password:                 string(hashedPassword),
		Telephone:                input.Telephone,
		Adresse:                  input.Adresse,
		Ville:                    input.Ville,
		Role:                     input.Role,
		EstAnimateurFormateur:    input.EstAnimateurFormateur,
		EstModerateur:            input.EstModerateur,
		EstResponsableValidation: input.EstResponsableValidation,
		EstServiceCheck:          input.EstServiceCheck,
		StatutCompte:             input.StatutCompte,
		DateFinSuspension:        input.DateFinSuspension,
		TutorielVu:               input.TutorielVu,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer l'utilisateur"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
func GetUsers(c *gin.Context) {
	var users []models.User
	query := database.DB

	role := c.Query("role")
	if role != "" {
		query = query.Where("role = ?", role)
	}

	query.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur introuvable"})
		return
	}

	c.JSON(http.StatusOK, user)
}


func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur introuvable"})
		return
	}

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requête invalide"})
		return
	}

	delete(input, "password")
	delete(input, "id")
	delete(input, "ID")

	database.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur introuvable"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé"})
}
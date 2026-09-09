package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"meridian/back/api/models"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("meridian.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données : ", err)
	}

	database.AutoMigrate(&models.User{},&models.Prestation{},&models.CategoriePrestation{})

	DB = database
}
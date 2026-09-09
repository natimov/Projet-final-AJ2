package models
import "gorm.io/gorm"

type CategoriePrestation struct {
	gorm.Model
	Libelle string `json:"libelle"`
}
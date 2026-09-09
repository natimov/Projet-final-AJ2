package models
import "gorm.io/gorm"

type Prestation struct {
	gorm.Model
	Nom            string  `json:"nom"`
	Description     string  `json:"description"`
	Prix          float64 `json:"prix"`
	Duree         int     `json:"duree"`
	NombrePlaces     int     `json:"nombre_places"`
	Date          string  `json:"date"`
	Lieu             string `json:"lieu"`
	ImageIllustration string `json:"image_illustration"`
	Statut           string  `json:"statut" gorm:"default:brouillon"`
}

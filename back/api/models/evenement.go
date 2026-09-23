package models

import "gorm.io/gorm"

type Evenement struct {
	gorm.Model
	Type      string `json:"type"`
	Titre     string `json:"titre"`
	DateDebut string `json:"date_debut"`
	DateFin   string `json:"date_fin"`
	Lieu      string `json:"lieu"`
}
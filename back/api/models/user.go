package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nom                      string `json:"nom"`
	Prenom                   string `json:"prenom"`
	Email                    string `json:"email" gorm:"unique"`
	Password                 string `json:"-"`
	Telephone                string `json:"telephone"`
	Adresse                  string `json:"adresse"`
	Ville                    string `json:"ville"`
	Role                     string `json:"role" gorm:"default:particulier"`
	EstAnimateurFormateur    bool   `json:"est_animateur_formateur" gorm:"default:false"`
	EstModerateur            bool   `json:"est_moderateur" gorm:"default:false"`
	EstResponsableValidation bool   `json:"est_responsable_validation" gorm:"default:false"`
	EstServiceCheck          bool   `json:"est_service_check" gorm:"default:false"`
	StatutCompte             string `json:"statut_compte" gorm:"default:actif"`
	DateFinSuspension        string `json:"date_fin_suspension"`
	TutorielVu               bool   `json:"tutoriel_vu" gorm:"default:false"`
}
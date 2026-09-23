package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret []byte

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erreur : fichier .env introuvable")
	}

	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

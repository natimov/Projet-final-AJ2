# UpcycleConnect

Projet composé d'une API en Go avec Gin et GORM, d'une base SQLite et d'une interface Vue.js.

## Prérequis

- Go (la version du module est indiquée dans `back/api/go.mod`)
- Node.js 22.18 ou plus récent et npm (voir `front/package.json`)

## Configuration et lancement

### Backend

Dans `back/api`, créez un fichier `.env` contenant une clé JWT :

```env
JWT_SECRET=une-cle-secrete-personnelle
```

Installez et lancez l'API depuis ce dossier :

```powershell
go mod download
go run .
```

L'API écoute sur `http://localhost:8080`. SQLite crée ou utilise `back/api/meridian.db`. Les tables sont préparées au démarrage par GORM.

### Frontend

Dans un autre terminal, depuis `front` :

```powershell
npm install
npm run dev
```

Vite affiche l'adresse locale, généralement `http://localhost:5173`.

## Tests backend

Depuis `back/api` :

```powershell
go test ./...
```

Les tests d'intégration utilisent des fichiers SQLite temporaires et ne réutilisent pas `meridian.db`.

## Structure

- `back/api/main.go` et `routes.go` : démarrage et routes de l'API
- `back/api/handlers` : traitement des requêtes
- `back/api/middlewares` : authentification, contrôle admin et CORS
- `back/api/models` : modèles GORM
- `back/api/database` : connexion SQLite et migrations
- `front/src` : application Vue.js

## Authentification

`POST /register` crée un compte particulier et `POST /login` renvoie un JWT. Les routes de gestion des utilisateurs, prestations, catégories et événements demandent un JWT d'un compte ayant le rôle `admin`. Envoyez-le dans l'en-tête `Authorization: Bearer <token>`.

## Git et fichiers locaux

Travaillez sur la branche de tâche convenue. Ne partagez pas `.env` et ne l'ajoutez pas à Git. La base `meridian.db` contient les données locales de développement.

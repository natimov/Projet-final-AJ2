# Projet-final-AJ2 — Meridian Tech

Back en Go natif (Gin, GORM, golang-jwt, bcrypt). Front en Vue.js.

## Prérequis à installer

- Node.js v24 (LTS) — nodejs.org
- Go 1.27 — go.dev/dl
- Postman (pour tester l'API) — postman.com/downloads

Vérifier les versions après installation :

```
node -v      # doit afficher v24.x
go version   # doit afficher go1.27.x
```

## Cloner le projet

```
git clone https://github.com/natimov/Projet-final-AJ2.git
cd Projet-final-AJ2
git checkout dev
```

## Installer le back (back/api)

```
cd back/api
copy .env.exemple .env
```

Ouvrir le fichier `.env` créé et remplacer la valeur de `JWT_SECRET` par une vraie clé secrète (n'importe quelle chaîne de caractères longue suffit en local).

```
go run .
```

Les dépendances (Gin, GORM, golang-jwt, bcrypt...) se téléchargent automatiquement au premier lancement.

## Installer le front (front)

```
cd ../front
npm install --legacy-peer-deps
```

Le `--legacy-peer-deps` est nécessaire à cause d'un conflit de versions entre `oxlint` et `eslint-plugin-oxlint`.

## Lancer le projet (2 terminaux séparés)

```
cd back/api
go run .
```

```
cd front
npm run dev
```

- API : http://localhost:8080
- Front : http://localhost:5173

## Conventions

- Jamais de push direct sur `main`
- Respecter la convention de nomage des commit (voir le PDF )
- Créer une branche par feature (ou groupe de feature )
- Merge sur dev uniquement si la feature est approuvée

  

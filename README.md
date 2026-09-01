# Projet-final-AJ2 — Meridian Tech

## Prérequis à installer

- Node.js v24 (LTS) 
- PHP 8.4 + Composer, via Laravel Herd 




```
node -v      
php -v       
composer -V
```

## Cloner le projet

```
git clone https://github.com/natimov/Projet-final-AJ2.git
cd Projet-final-AJ2
git checkout dev
```

## Installer le back (api)

```
cd api
composer install
copy .env.example .env
php artisan key:generate
php artisan migrate
```

## Installer le front (front)

```
cd ../front
npm install --legacy-peer-deps
```

Le `--legacy-peer-deps` est nécessaire à cause d'un conflit de versions entre `oxlint` et `eslint-plugin-oxlint`.

## Lancer le projet (2 pwshell séparés)

```
cd api
php artisan serve
```

```
cd front
npm run dev
```

- API : http://127.0.0.1:8000
- Front : http://localhost:5173


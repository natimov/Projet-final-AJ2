<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="header-top">
      <img alt="Logo Meridian" class="logo" src="@/assets/meridian-logo.png" />
      <div class="brand">
        <h1>Meridian</h1>
        <p class="tagline">Bienvenue sur l'application Meridian</p>
      </div>

      <nav>
        <RouterLink to="/">Accueil</RouterLink>
        <RouterLink to="/login">Connexion</RouterLink>
        <RouterLink to="/register">Créer un compte</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  background-color: var(--color-primary);
  padding: 1.5rem 2rem;
}
.logo {
  height: 120px;
  filter: brightness(0) invert(1);
  margin-right: 1rem;
}

.header-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  max-width: 1280px;
  margin: 0 auto;
}

.brand h1 {
  font-family: var(--font-heading);
  font-weight: 700;
  font-size: 1.8rem;
  color: #ffffff;
  margin: 0;
}

.tagline {
  font-family: var(--font-body);
  font-size: 0.9rem;
  color: var(--color-neutral);
  margin-top: 0.25rem;
}

nav {
  display: flex;
  gap: 1.5rem;
  padding-top: 0.4rem;
}

nav a {
  font-family: var(--font-body);
  color: #ffffff;
  text-decoration: none;
  font-size: 0.95rem;
}

nav a:hover,
nav a.router-link-exact-active {
  color: var(--color-accent);
}
</style>
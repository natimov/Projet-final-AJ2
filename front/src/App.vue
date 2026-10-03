<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from './stores/auth'

const authStore = useAuthStore()
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { onClickOutside } from '@vueuse/core'

const router = useRouter()
const showMenu = ref(false)
const profileRef = ref(null)

function toggleMenu() {
  showMenu.value = !showMenu.value
}

function handleLogout() {
  authStore.logout()
  showMenu.value = false
  router.push('/login')
}

onClickOutside(profileRef, () => {
  showMenu.value = false
})
</script>

<template>
  <header>
    <div class="header-top">
      <div class="brand">
        <h1>Upcycle Connect</h1>
        <p class="tagline">Bienvenue sur l'application Upcycle Connect</p>
      </div>

      <nav>
        <RouterLink to="/">Accueil</RouterLink>

        <template v-if="!authStore.user">
          <RouterLink to="/login">Connexion</RouterLink>
          <RouterLink to="/register">Inscription</RouterLink>
        </template>

        <div v-else class="profile" ref="profileRef" @click="toggleMenu">
          <svg viewBox="0 0 24 24" width="28" height="28" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="8" r="4" />
            <path d="M4 20c0-4 4-6 8-6s8 2 8 6" />
          </svg>
          <span>{{ authStore.user.prenom }}</span>

          <div v-if="showMenu" class="dropdown">
            <button @click="handleLogout">Se déconnecter</button>
          </div>
        </div>
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
.profile {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #ffffff;
  font-family: var(--font-body);
  font-size: 0.85rem;
  gap: 0.2rem;
  position: relative;
  cursor: pointer;
}
.dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 0.5rem;
  background-color: #ffffff;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  overflow: hidden;
}
.dropdown button {
  display: block;
  width: 100%;
  padding: 0.6rem 1.2rem;
  border: none;
  background: none;
  color: var(--color-text);
  font-family: var(--font-body);
  font-size: 0.85rem;
  cursor: pointer;
  white-space: nowrap;
}
</style>
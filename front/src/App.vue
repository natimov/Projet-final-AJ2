<script setup>
import { computed, ref, watch } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import { useI18n } from 'vue-i18n'

const { t, locale, availableLocales } = useI18n({ useScope: 'global' })

const languages = computed(() => {
  const names = new Intl.DisplayNames([locale.value], { type: 'language' })

  return availableLocales.map((code) => ({
    code,
    name: names.of(code) || code,
  }))
})

watch(
  locale,
  (language) => {
    localStorage.setItem('language', language)
    document.documentElement.lang = language
    document.title = t('common.siteName')
  },
  { immediate: true },
)
import { useAuthStore } from './stores/auth'

const authStore = useAuthStore()
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
        <RouterLink class="brand-link" to="/">{{ t('common.siteName') }}</RouterLink>
        <p class="tagline">{{ t('common.tagline') }}</p>
      </div>

      <nav>
        <RouterLink to="/">{{ t('common.home') }}</RouterLink>
        <RouterLink to="/espace-particulier">{{ t('common.privateSpace') }}</RouterLink>
        <RouterLink to="/espace-professionnel">{{ t('common.professionalSpace') }}</RouterLink>
        <RouterLink to="/espace-salarie">{{ t('common.employeeSpace') }}</RouterLink>
        <RouterLink to="/admin">{{ t('common.admin') }}</RouterLink>
      </nav>

      <div class="header-actions">
        <label class="language-control">
          <span>{{ t('common.language') }}</span>
          <select v-model="locale">
            <option v-for="language in languages" :key="language.code" :value="language.code">
              {{ language.name }}
            </option>
          </select>
        </label>
                <template v-if="!authStore.user">
          <RouterLink to="/login">{{ t('common.login') }}</RouterLink>
          <RouterLink class="register-link" to="/register">{{ t('common.register') }}</RouterLink>
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
      </div>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  background-color: #244b38;
  padding: 1.1rem clamp(1rem, 3vw, 3rem);
}

.header-top {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  gap: 1.5rem;
}

.brand {
  justify-self: start;
}

.brand-link {
  color: #ffffff;
  font-family: var(--font-heading);
  font-weight: 700;
  font-size: 1.5rem;
  text-decoration: none;
}

.tagline {
  font-family: var(--font-body);
  font-size: 0.9rem;
  color: #dce6d7;
  margin-top: 0.1rem;
}

nav {
  display: flex;
  justify-self: center;
  gap: 2rem;
}

nav a {
  font-family: var(--font-body);
  color: #ffffff;
  text-decoration: none;
  font-size: 0.95rem;
  white-space: nowrap;
}

nav a:hover,
nav a.router-link-exact-active,
.brand-link:hover {
  color: #d4e49e;
}

nav a:focus-visible,
.brand-link:focus-visible {
  outline: 2px solid #d4e49e;
  outline-offset: 4px;
}

.language-control {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #ffffff;
  font-size: 0.85rem;
  white-space: nowrap;
}

.header-actions {
  display: flex;
  align-items: center;
  justify-self: end;
  gap: 1.25rem;
}

.header-actions > a {
  color: #ffffff;
  font-family: var(--font-body);
  font-size: 0.95rem;
  text-decoration: none;
  white-space: nowrap;
}

.header-actions > a:hover {
  color: #d4e49e;
}

.register-link {
  flex-shrink: 0;
  padding: 0.55rem 0.9rem;
  border: 1px solid #dce6d7;
  border-radius: 24px;
}

.language-control select {
  padding: 0.35rem 0.5rem;
  border: 1px solid #dce6d7;
  border-radius: 4px;
  color: #244b38;
  background: #ffffff;
}

@media (max-width: 1150px) {
  .header-top {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: 1rem;
  }

  .brand {
    order: 1;
  }

  .header-actions {
    order: 2;
    flex-wrap: wrap;
    gap: 0.75rem;
  }

  nav {
    order: 3;
    flex: 1 0 100%;
    justify-self: auto;
    flex-wrap: wrap;
    justify-content: center;
    gap: 0.75rem 1.5rem;
  }
}

@media (max-width: 700px) {
  .header-top {
    align-items: stretch;
    flex-direction: column;
    gap: 0.9rem;
  }

  .brand {
    align-self: flex-start;
  }

  nav {
    order: 2;
    flex: none;
    justify-self: auto;
    justify-content: flex-start;
    flex-wrap: wrap;
    gap: 0.65rem 1rem;
  }

  .header-actions {
    order: 3;
    justify-self: auto;
    flex-wrap: wrap;
    gap: 0.75rem;
  }
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

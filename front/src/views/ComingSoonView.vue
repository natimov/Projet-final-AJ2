<script setup>
import { RouterLink, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'

const route = useRoute()
const { t } = useI18n()
</script>

<template>
  <main class="coming-soon">
    <p class="eyebrow">{{ t('spaces.label') }}</p>
    <h1>{{ t(route.meta.titleKey) }}</h1>
    <p class="message">{{ t('spaces.comingSoon') }}</p>
    <RouterLink class="home-link" to="/">{{ t('common.home') }}</RouterLink>
  </main>
</template>

<style scoped>
.coming-soon {
  display: flex;
  min-height: 60vh;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 3rem 1.5rem;
  text-align: center;
}

.eyebrow {
  color: #607a52;
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.12em;
}

h1 {
  margin-top: 0.75rem;
  color: #244b38;
  font-family: var(--font-heading);
  font-size: clamp(2rem, 5vw, 3.5rem);
}

.message {
  margin-top: 0.75rem;
  color: #637168;
  font-size: 1.1rem;
}

.home-link {
  margin-top: 1.5rem;
  padding: 0.65rem 1rem;
  border-radius: 24px;
  color: #ffffff;
  background: #244b38;
  text-decoration: none;
}

.home-link:hover {
  background: #466d4f;
}
</style>

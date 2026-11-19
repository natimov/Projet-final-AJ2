
<script setup>
import { computed, nextTick, ref, watch } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAuthStore } from '../stores/auth'
import ParticularRequests from '../components/particulier/ParticularRequests.vue'
import ParticularCatalogue from '../components/particulier/ParticularCatalogue.vue'
import ParticularAdvice from '../components/particulier/ParticularAdvice.vue'
import ParticularPlanning from '../components/particulier/ParticularPlanning.vue'
import ParticularTutorial from '../components/particulier/ParticularTutorial.vue'
import { useLocalValue } from '../components/particulier/local.js'
import '../assets/particulier.css'

const { t, n } = useI18n()
const auth = useAuthStore()
const route = useRoute()
const router = useRouter()
const canDemo = import.meta.env.DEV
const demo = computed(() => canDemo && route.query.demo === '1')
const sections = ['overview', 'listings', 'deposits', 'advice', 'catalogue', 'planning', 'score', 'profile']
const section = computed(() => sections.includes(route.query.section) ? route.query.section : 'overview')
const account = computed(() => auth.user || {})
const accountKey = computed(() => account.value.ID ?? account.value.id ?? account.value.email)
const workspaceKey = computed(() => [accountKey.value, demo.value ? 'demo' : 'drafts'].join(':'))
const { value: bookings, save: saveBookings } = useLocalValue(workspaceKey, 'bookings', [])
const { value: tutorialSeen, save: saveTutorial } = useLocalValue(accountKey, 'tutorial', false)
const tutorial = ref(null)
const content = ref(null)
const storageError = ref(false)

const profile = computed(() => [
  ['nameLabel', [account.value.prenom, account.value.nom].filter(Boolean).join(' ')],
  ['emailLabel', account.value.email],
  ['phoneLabel', account.value.telephone],
  ['addressLabel', account.value.adresse],
  ['cityLabel', account.value.ville],
  ['statusLabel', t('particular.accountStates.' + (
    ['actif', 'suspendu', 'banni'].includes(account.value.statut_compte)
      ? account.value.statut_compte
      : 'unknown'
  ))],
])
const impact = [
  ['scoreValue', 120], ['objects', 12], ['waste', 21], ['carbon', 36],
]

function finishTutorial() {
  storageError.value = !saveTutorial(true)
}

function setDemo() {
  const query = { ...route.query, demo: demo.value ? undefined : '1' }
  router.replace({ query })
}

watch(section, async () => {
  await nextTick()
  content.value?.focus()
})
</script>

<template>
  <main class="particular-page">
    <div class="particular-shell">
      <header class="particular-heading">
        <div>
          <p>{{ t('spaces.particular') }}</p>
          <h1>{{ t('particular.greeting', { name: account.prenom || t('particular.member') }) }}</h1>
        </div>
        <button type="button" class="p-button p-secondary" @click="tutorial?.open()">
          {{ t('particular.tutorial.open') }}
        </button>
      </header>

      <div v-if="canDemo" class="p-demo">
        <p v-if="demo"><strong>{{ t('particular.demoTitle') }}</strong> {{ t('particular.demoText') }}</p>
        <button type="button" class="p-button p-secondary" :aria-pressed="demo" @click="setDemo">
          {{ t(demo ? 'particular.exitDemo' : 'particular.startDemo') }}
        </button>
      </div>

      <div class="particular-layout">
        <nav class="p-navigation" :aria-label="t('spaces.particular')">
          <RouterLink v-for="item in sections" :key="item"
            :to="{ name: 'particular-space', query: { section: item === 'overview' ? undefined : item, demo: demo ? '1' : undefined } }"
            :aria-current="section === item ? 'page' : null"
            :class="{ 'p-current': section === item }">
            {{ t('particular.sections.' + item) }}
          </RouterLink>
        </nav>

        <section ref="content" tabindex="-1" class="p-section" aria-labelledby="particular-section-heading">
          <h2 id="particular-section-heading">{{ t('particular.sections.' + section) }}</h2>

          <ul v-if="section === 'overview'" class="p-overview p-grid">
            <li v-for="item in sections.slice(1)" :key="item" class="p-card">
              <h3>
                <RouterLink :to="{ name: 'particular-space', query: { section: item, demo: demo ? '1' : undefined } }">
                  {{ t('particular.sections.' + item) }}
                </RouterLink>
              </h3>
              <p>{{ t('particular.overview.' + item) }}</p>
            </li>
          </ul>

          <ParticularRequests
            v-else-if="section === 'listings' || section === 'deposits'"
            :key="workspaceKey + ':' + section" :kind="section" :scope="workspaceKey" :demo="demo"
          />
          <ParticularAdvice v-else-if="section === 'advice'" :key="workspaceKey" :scope="workspaceKey" :demo="demo" />
          <ParticularCatalogue
            v-else-if="section === 'catalogue'"
            :demo="demo" :token="auth.token" :bookings="bookings"
            :save-bookings="saveBookings"
          />
          <ParticularPlanning v-else-if="section === 'planning'" :demo="demo" :bookings="bookings" />

          <template v-else-if="section === 'score'">
            <p>{{ t('particular.overview.score') }}</p>
            <p v-if="!demo" class="p-notice">{{ t('particular.scoreUnavailable') }}</p>
            <dl class="p-grid p-metrics">
              <div v-for="[key, value] in impact" :key="key" class="p-card">
                <dt>{{ t('particular.metrics.' + key) }}</dt>
                <dd>{{ demo ? n(value) : '—' }}</dd>
              </div>
              <div class="p-card">
                <dt>{{ t('particular.metrics.badge') }}</dt>
                <dd>{{ demo ? t('particular.metrics.starter') : '—' }}</dd>
              </div>
            </dl>
          </template>

          <dl v-else class="p-grid p-profile">
            <div v-for="[label, value] in profile" :key="label" class="p-card">
              <dt>{{ t('particular.' + label) }}</dt>
              <dd>{{ value || '—' }}</dd>
            </div>
          </dl>
        </section>
      </div>

      <p v-if="storageError" role="alert" class="p-notice">{{ t('particular.storageError') }}</p>
      <ParticularTutorial
        :key="accountKey" ref="tutorial"
        :first-visit="!tutorialSeen && !account.tutoriel_vu"
        @finished="finishTutorial"
      />
    </div>
  </main>
</template>

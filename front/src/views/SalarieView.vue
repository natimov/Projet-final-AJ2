<script setup>
import { computed, nextTick, onMounted, ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAuthStore } from '../stores/auth'
import FormationsSalarie from '../components/salarie/FormationsSalarie.vue'
import PlanningSalarie from '../components/salarie/PlanningSalarie.vue'
import ConseilsSalarie from '../components/salarie/ConseilsSalarie.vue'
import ModerationSalarie from '../components/salarie/ModerationSalarie.vue'
import { useEmployeeDrafts } from '../components/salarie/employee'
import '../assets/salarie.css'

const { t } = useI18n()
const route = useRoute()
const auth = useAuthStore()
const authorized = computed(() => Boolean(auth.token && auth.user))
const sections = ['overview', 'training', 'planning', 'advice', 'moderation']
const section = computed(() => sections.includes(route.query.section) ? route.query.section : 'overview')
const { data: drafts, storageError: draftError, save: saveDrafts } = useEmployeeDrafts(() => authorized.value ? (auth.user?.ID ?? auth.user?.id ?? auth.user?.email) : null)
const api = ref({ trainings: [], events: [], loading: false, error: '' })
const refresh = ref(0)
const content = ref(null)
const tutorial = ref(null)
const tutorialStep = ref(0)
const tutorSteps = sections
const author = computed(() => [auth.user?.prenom, auth.user?.nom].filter(Boolean).join(' '))
const displayName = computed(() => auth.user?.prenom || t('employee.member'))
const links = sections.filter((item) => item !== 'overview')
const trainings = computed(() => [...api.value.trainings, ...drafts.value.trainings])
const events = computed(() => [...api.value.events, ...drafts.value.events])

function saveItems(name, items) {
  return saveDrafts(name, items.filter((item) => item.local))
}

function sectionLink(item) {
  return { name: 'employee-space', query: { section: item === 'overview' ? undefined : item } }
}

function stateName(value) {
  return ({ brouillon: 'draft', en_attente: 'pending', en_attente_validation: 'pending', valide: 'validated', validee: 'validated', refuse: 'refused', refusee: 'refused', publie: 'published', publiee: 'published', termine: 'completed', terminee: 'completed', annule: 'cancelled', annulee: 'cancelled' })[value] || 'unknown'
}

function typeName(value) {
  const type = String(value || '').normalize('NFD').replace(/\p{Diacritic}/gu, '').toLowerCase()
  return ({ atelier: 'workshop', formation: 'training', evenement: 'event', reunion: 'meeting', conference: 'conference', permanence: 'office', conge: 'leave' })[type] || 'activity'
}

async function readList(path, signal) {
  const response = await fetch('/api/' + path, { headers: { Authorization: 'Bearer ' + auth.token }, signal })
  if (response.status === 401 || response.status === 403) throw new Error('restricted')
  if (!response.ok) throw new Error('unavailable')
  const result = await response.json()
  if (result !== null && !Array.isArray(result)) throw new Error('unavailable')
  return result || []
}

watch([authorized, () => auth.token, refresh], async ([allowed, token], _, onCleanup) => {
  api.value = { trainings: [], events: [], loading: false, error: '' }
  if (!allowed || !token) return
  api.value.loading = true
  const controller = new AbortController()
  onCleanup(() => controller.abort())
  try {
    const [trainings, events] = await Promise.all([readList('prestations', controller.signal), readList('evenements', controller.signal)])
    if (controller.signal.aborted) return
    api.value.trainings = trainings.map((item) => ({ id: 'training-' + item.ID, type: typeName(item.categorie?.libelle), title: item.nom, description: item.description, start: item.date, duration: item.duree || 0, capacity: item.nombre_places || 0, price: item.prix, location: item.lieu, objectives: '', status: stateName(item.statut), rejectionReason: item.motif_refus || '', participants: [], waiting: [] }))
    api.value.events = events.map((item) => ({ id: 'event-' + item.ID, type: typeName(item.type), title: item.titre, start: item.date_debut, end: item.date_fin, location: item.lieu }))
  } catch (error) {
    if (!controller.signal.aborted) api.value.error = error.message === 'restricted' ? 'restricted' : 'unavailable'
  } finally {
    if (!controller.signal.aborted) api.value.loading = false
  }
}, { immediate: true })

watch(section, async () => {
  await nextTick()
  content.value?.focus()
})

function openTutorial() {
  tutorialStep.value = 0
  if (tutorial.value && !tutorial.value.open) tutorial.value.showModal()
}

function closeTutorial() {
  tutorial.value?.close()
  try {
    localStorage.setItem('upcycleconnect:employee:tutorial:' + auth.user?.ID, '1')
  } catch {
    return
  }
}

function nextTutorial() {
  if (tutorialStep.value === tutorSteps.length - 1) closeTutorial()
  else tutorialStep.value++
}

onMounted(() => {
  if (!authorized.value) return
  try {
    const seen = localStorage.getItem('upcycleconnect:employee:tutorial:' + auth.user?.ID)
    if (!seen && !auth.user?.tutoriel_vu) openTutorial()
  } catch {
    return
  }
})
</script>

<template>
  <main v-if="!authorized" class="employee-page">
    <div class="employee-shell">
      <h1 class="employee-access-title">{{ t('spaces.employee') }}</h1>
      <div class="e-panel">
        <p>{{ t('employee.accessDenied') }}</p>
        <div class="e-actions">
          <RouterLink class="e-button" :to="{ path: '/login', query: { redirect: '/espace-salarie' } }">{{ t('employee.employeeLogin') }}</RouterLink>
        </div>
      </div>
    </div>
  </main>
  <main v-else class="employee-page">
    <div class="employee-shell">
      <header class="employee-heading"><div><p>{{ t('spaces.employee') }}</p><h1>{{ t('employee.greeting', { name: displayName }) }}</h1></div><button class="e-button e-secondary" type="button" @click="openTutorial">{{ t('employee.tutorial.open') }}</button></header>
      <p v-if="draftError" class="e-error" role="alert">{{ t('employee.storageError') }}</p>
      <div class="employee-layout">
        <nav class="employee-nav" :aria-label="t('spaces.employee')"><RouterLink v-for="item in sections" :key="item" :to="sectionLink(item)" :aria-current="section === item ? 'page' : null" :class="{ active: section === item }">{{ t('employee.sections.' + item) }}</RouterLink></nav>
        <section ref="content" class="employee-content" tabindex="-1" aria-labelledby="employee-section-title">
          <h2 id="employee-section-title">{{ t('employee.sections.' + section) }}</h2>
          <template v-if="section === 'overview'"><p>{{ t('employee.overview') }}</p><ul class="e-overview"><li v-for="item in links" :key="item" class="e-panel"><h3><RouterLink :to="sectionLink(item)">{{ t('employee.sections.' + item) }}</RouterLink></h3><p>{{ t('employee.descriptions.' + item) }}</p></li></ul></template>
          <template v-else>
            <p v-if="api.loading && ['training', 'planning'].includes(section)" class="e-notice" role="status">{{ t('employee.loading') }}</p>
            <div v-else-if="api.error && ['training', 'planning'].includes(section)" class="e-notice" role="status"><p>{{ t('employee.api.' + api.error) }}</p><button class="e-button e-secondary" type="button" @click="refresh++">{{ t('employee.retry') }}</button></div>
            <p v-else-if="['training', 'planning'].includes(section)" class="e-notice">{{ t('employee.api.generalList') }}</p>
            <p v-if="section === 'advice'" class="e-notice" role="status">{{ t('employee.api.unavailableSection') }}</p>
            <FormationsSalarie v-if="section === 'training'" :items="trainings" :events="events" :remote-unavailable="Boolean(api.error || api.loading)" :save-items="(items) => saveItems('trainings', items)" />
            <PlanningSalarie v-if="section === 'planning'" :events="events" :trainings="trainings" :remote-unavailable="Boolean(api.error || api.loading)" :save-items="(items) => saveItems('events', items)" />
            <ConseilsSalarie v-if="section === 'advice'" :items="drafts.advice" :author="author" :save-items="(items) => saveItems('advice', items)" />
            <ModerationSalarie v-if="section === 'moderation'" />
          </template>
        </section>
      </div>
      <dialog ref="tutorial" class="employee-tutorial" aria-labelledby="employee-tutorial-title" @cancel.prevent="closeTutorial"><p>{{ t('employee.tutorial.progress', { current: tutorialStep + 1, total: tutorSteps.length }) }}</p><h2 id="employee-tutorial-title">{{ t('employee.sections.' + tutorSteps[tutorialStep]) }}</h2><p>{{ t('employee.tutorial.steps.' + tutorSteps[tutorialStep]) }}</p><div class="e-actions"><button class="e-button e-secondary" type="button" autofocus @click="closeTutorial">{{ t('employee.tutorial.skip') }}</button><button v-if="tutorialStep" class="e-button e-secondary" type="button" @click="tutorialStep--">{{ t('employee.tutorial.previous') }}</button><button class="e-button" type="button" @click="nextTutorial">{{ t(tutorialStep === tutorSteps.length - 1 ? 'employee.tutorial.finish' : 'employee.tutorial.next') }}</button></div></dialog>
    </div>
  </main>
</template>

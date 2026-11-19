<script setup>
import { computed, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { dateInDays } from './local'

const props = defineProps({
  demo: Boolean,
  token: { type: String, default: null },
  bookings: { type: Array, required: true },
  saveBookings: { type: Function, required: true },
})
const { t, n, d } = useI18n()
const kinds = ['workshop', 'training', 'event', 'service']
const samples = kinds.map((kind, index) => ({
  id: 'demo-' + kind, kind, titleKey: kind + 'Title', descriptionKey: kind + 'Text',
  startsAt: dateInDays(2 + index * 3) + 'T14:00:00',
  duration: [90, 120, 60, 45][index],
  price: [0, 24, 0, 12][index], seats: [8, 12, 30, 3][index],
  location: ['Paris 11e', 'Paris 13e', 'Montreuil', 'Ivry-sur-Seine'][index],
  host: ['Camille', 'Nora', 'Alex', 'Lou'][index],
}))

const activities = ref([])
const loadState = ref('loading')
const warning = ref('')
const reload = ref(0)
const search = ref('')
const kind = ref('')
const feedback = ref('')
const storageError = ref(false)
const paymentPreview = ref(null)

function title(activity) {
  return activity.titleKey ? t('particular.catalogue.samples.' + activity.titleKey) : activity.title
}

const visible = computed(() => {
  const query = search.value.trim().toLocaleLowerCase()
  const list = props.demo ? samples.map((activity) => props.bookings.find(({ id }) => id === activity.id) ?? activity) : activities.value
  return list.filter((activity) => (
    (!kind.value || activity.kind === kind.value)
    && [title(activity), activity.location, activity.category].some((value) => value?.toLocaleLowerCase().includes(query))
  ))
})

async function read(path, token, signal) {
  const response = await fetch('/api/' + path, { headers: { Authorization: 'Bearer ' + token }, signal })
  if (response.status === 401 || response.status === 403) return { state: 'restricted' }
  if (!response.ok) throw new Error('HTTP ' + response.status)
  const data = await response.json()
  if (data !== null && !Array.isArray(data)) throw new Error('Expected a list')
  return { state: 'ready', data: data ?? [] }
}

function label(value) {
  return value?.normalize('NFD').replace(/\p{Diacritic}/gu, '').toLowerCase().trim()
}

function offeringKind(category) {
  const value = label(category) || ''
  if (value.includes('atelier')) return 'workshop'
  if (value.includes('formation')) return 'training'
  return 'service'
}

watch([() => props.demo, () => props.token, reload], async ([demo, token], _, onCleanup) => {
  feedback.value = ''
  storageError.value = false
  paymentPreview.value = null
  activities.value = []
  warning.value = ''
  if (demo) {
    loadState.value = 'ready'
    return
  }
  if (!token) {
    loadState.value = 'restricted'
    return
  }
  loadState.value = 'loading'
  const controller = new AbortController()
  onCleanup(() => controller.abort())

  const results = await Promise.allSettled(['prestations', 'evenements', 'categories-prestation']
    .map((path) => read(path, token, controller.signal)))
  if (controller.signal.aborted) return

  const [offerings, events, categories] = results.map((result) => (
    result.status === 'fulfilled' ? result.value : { state: 'error' }
  ))
  const categoryNames = new Map((categories.data || []).map((item) => [item.ID, item.libelle]))
  const offeringsList = (offerings.data || [])
    .filter((item) => ['publie', 'publiee', 'published'].includes(label(item.statut)))
    .map((item) => {
      const category = item.categorie?.libelle || categoryNames.get(item.categorie_id)
      return {
        id: 'offering-' + item.ID, title: item.nom, description: item.description, category,
        kind: offeringKind(category), startsAt: item.date, location: item.lieu, duration: item.duree,
        seats: item.nombre_places, price: item.prix,
      }
    })

  activities.value = [...offeringsList, ...(events.data || []).map((item) => ({
    id: 'event-' + item.ID, kind: 'event', title: item.titre, category: item.type,
    startsAt: item.date_debut, endsAt: item.date_fin, location: item.lieu,
  }))]
  const states = [offerings.state, events.state]
  const blocked = states.includes('error') ? 'error' : 'restricted'
  loadState.value = states.includes('ready') ? 'ready' : blocked
  warning.value = states.includes('ready') && states.some((state) => state !== 'ready') ? blocked : ''
}, { immediate: true })

function dateLabel(value) {
  if (!value) return '—'
  const parsed = new Date(value.length === 10 ? value + 'T12:00:00' : value)
  if (Number.isNaN(parsed.getTime())) return value || '—'
  return d(parsed, value.length === 10 ? { dateStyle: 'medium' } : { dateStyle: 'medium', timeStyle: 'short' })
}

function isBooked(id) {
  return props.bookings.some((booking) => booking.id === id)
}

function startBooking(activity) {
  if (!props.demo) return
  if (!isBooked(activity.id) && activity.price > 0) {
    paymentPreview.value = activity.id
    return
  }
  toggleBooking(activity)
}

function toggleBooking(activity) {
  if (!props.demo) return
  const booked = isBooked(activity.id)
  const next = booked
    ? props.bookings.filter((booking) => booking.id !== activity.id)
    : [...props.bookings, activity]

  feedback.value = ''
  storageError.value = !props.saveBookings(next)
  if (!storageError.value) {
    paymentPreview.value = null
    feedback.value = booked ? 'demoCancelled' : 'demoBooked'
  }
}
</script>

<template>
  <p>{{ t('particular.overview.catalogue') }}</p>
  <p v-if="loadState === 'loading'" class="p-notice" role="status">{{ t('particular.catalogue.loading') }}</p>
  <div v-else-if="loadState !== 'ready'" class="p-notice">
    <p role="status">{{ t('particular.catalogue.' + loadState) }}</p>
    <button v-if="loadState === 'error'" class="p-button p-secondary" type="button" @click="reload++">
      {{ t('particular.catalogue.retry') }}
    </button>
  </div>

  <template v-else>
    <p v-if="warning" class="p-notice" role="status">{{ t('particular.catalogue.' + warning) }}</p>
    <div class="p-filters">
      <label>{{ t('particular.search') }} <input v-model="search" type="search" /></label>
      <label>{{ t('particular.catalogue.type') }}
        <select v-model="kind">
          <option value="">{{ t('particular.all') }}</option>
          <option v-for="item in kinds" :value="item" :key="item">{{ t('particular.catalogue.kinds.' + item) }}</option>
        </select>
      </label>
    </div>

    <p v-if="storageError" class="p-notice" role="alert">{{ t('particular.storageError') }}</p>
    <p v-if="feedback" class="p-notice" role="status">{{ t('particular.catalogue.' + feedback) }}</p>
    <ul v-if="visible.length" class="p-rows">
      <li v-for="activity in visible" :key="activity.id">
        <details class="p-card p-details">
          <summary>
            <span class="p-chip">{{ t('particular.catalogue.kinds.' + activity.kind) }}</span>
            <strong>{{ title(activity) }}</strong>
            <span>{{ dateLabel(activity.startsAt) }} · {{ activity.location || '—' }}</span>
          </summary>

          <p v-if="activity.descriptionKey">{{ t('particular.catalogue.samples.' + activity.descriptionKey) }}</p>
          <p v-else-if="activity.description">{{ activity.description }}</p>
          <dl class="p-grid p-profile">
            <div v-if="activity.category"><dt>{{ t('particular.requests.category') }}</dt><dd>{{ activity.category }}</dd></div>
            <div v-if="activity.endsAt"><dt>{{ t('particular.catalogue.endsAt') }}</dt><dd>{{ dateLabel(activity.endsAt) }}</dd></div>
            <div v-if="activity.duration"><dt>{{ t('particular.catalogue.duration') }}</dt><dd>{{ t('particular.catalogue.minutes', { count: activity.duration }) }}</dd></div>
            <div v-if="activity.seats != null"><dt>{{ t('particular.catalogue.seats') }}</dt><dd>{{ n(activity.seats) }}</dd></div>
            <div v-if="activity.host"><dt>{{ t('particular.catalogue.host') }}</dt><dd>{{ activity.host }}</dd></div>
            <div v-if="activity.price != null"><dt>{{ t('particular.requests.price') }}</dt><dd>
              {{ activity.price > 0 ? n(activity.price, { style: 'currency', currency: 'EUR' }) : t('particular.catalogue.free') }}
            </dd></div>
          </dl>
          <template v-if="demo">
            <div v-if="paymentPreview === activity.id" class="p-notice">
              <h3>{{ t('particular.catalogue.paymentPreview') }}</h3>
              <p>{{ t('particular.catalogue.demoPayment', { total: n(activity.price, { style: 'currency', currency: 'EUR' }) }) }}</p>
              <div class="p-actions">
                <button type="button" class="p-button" @click="toggleBooking(activity)">
                  {{ t('particular.catalogue.confirmDemoPayment') }}
                </button>
                <button type="button" class="p-button p-secondary" @click="paymentPreview = null">{{ t('admin.cancel') }}</button>
              </div>
            </div>
            <div v-else class="p-actions">
              <button type="button" class="p-button" @click="startBooking(activity)">
                {{ t(isBooked(activity.id) ? 'particular.catalogue.cancelDemo' : 'particular.catalogue.bookDemo') }}
              </button>
            </div>
          </template>
          <p v-else class="p-notice">{{ t('particular.catalogue.bookingUnavailable') }}</p>
        </details>
      </li>
    </ul>
    <p v-else class="p-card">{{ t('particular.catalogue.empty') }}</p>
  </template>
</template>

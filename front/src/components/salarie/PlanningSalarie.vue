<script setup>
import { computed, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { dateAfter, dateLabel, endOfTraining, isConflict, newId, trainingStatus } from './employee'

const props = defineProps({ events: { type: Array, required: true }, trainings: { type: Array, required: true }, remoteUnavailable: Boolean, saveItems: { type: Function, required: true } })
const { t, locale } = useI18n()
const view = ref('week')
const day = ref(dateAfter(0).slice(0, 10))
const form = ref(null)
const error = ref('')
const reminderState = ref('inactive')
let reminderTimer
const types = ['training', 'workshop', 'meeting', 'conference', 'office', 'leave']
const period = computed(() => {
  const start = new Date((day.value || dateAfter(0).slice(0, 10)) + 'T00:00:00')
  const end = new Date(start)
  if (view.value === 'week') start.setDate(start.getDate() - (start.getDay() + 6) % 7)
  if (view.value === 'month') start.setDate(1)
  end.setTime(start.getTime())
  if (view.value === 'day') end.setDate(end.getDate() + 1)
  if (view.value === 'week') end.setDate(end.getDate() + 7)
  if (view.value === 'month') end.setMonth(end.getMonth() + 1)
  return { start, end }
})
const allEntries = computed(() => {
  const sessions = props.trainings.filter((item) => ['published', 'completed'].includes(trainingStatus(item))).map((item) => ({ ...item, end: endOfTraining(item), session: true }))
  return [...props.events, ...sessions].sort((a, b) => new Date(a.start) - new Date(b.start))
})
const entries = computed(() => allEntries.value.filter((item) => new Date(item.start) < period.value.end && new Date(item.end) > period.value.start))

function checkReminders() {
  const now = Date.now()
  for (const entry of allEntries.value) {
    const start = new Date(entry.start).getTime()
    if (start <= now || start > now + 86400000 || entry.local) continue
    const key = 'upcycleconnect:employee:reminder:' + entry.id + ':' + start
    try {
      if (sessionStorage.getItem(key)) continue
      new Notification(t('employee.planning.reminderTitle'), { body: entry.title + ' · ' + dateLabel(entry.start, locale.value) })
      sessionStorage.setItem(key, '1')
    } catch {
      reminderState.value = 'unavailable'
      clearInterval(reminderTimer)
      return
    }
  }
}

async function enableReminders() {
  if (!('Notification' in window) || !window.isSecureContext) {
    reminderState.value = 'unavailable'
    return
  }
  try {
    const permission = await Notification.requestPermission()
    if (permission !== 'granted') {
      reminderState.value = 'denied'
      return
    }
    clearInterval(reminderTimer)
    reminderState.value = 'active'
    checkReminders()
    if (reminderState.value === 'active') reminderTimer = setInterval(checkReminders, 60000)
  } catch {
    reminderState.value = 'unavailable'
  }
}

onUnmounted(() => clearInterval(reminderTimer))

function create() {
  error.value = ''
  const date = day.value || dateAfter(0).slice(0, 10)
  form.value = { id: newId(), local: true, type: 'meeting', title: '', start: date + 'T10:00', end: date + 'T11:00', location: '' }
}

function save() {
  const item = form.value
  const start = new Date(item.start)
  const end = new Date(item.end)
  if (!item.title.trim() || !item.location.trim() || Number.isNaN(start.getTime()) || Number.isNaN(end.getTime()) || end <= start) {
    error.value = 'formError'
    return
  }
  if (isConflict(start, end, props.events, props.trainings, item.id)) {
    error.value = 'conflict'
    return
  }
  if (!props.saveItems([...props.events.filter((entry) => entry.id !== item.id), { ...item }])) return
  form.value = null
  error.value = ''
}

function remove(item) {
  if (window.confirm(t('employee.confirmDeleteDraft'))) props.saveItems(props.events.filter((entry) => entry.id !== item.id))
}

function move(direction) {
  const date = new Date((day.value || dateAfter(0).slice(0, 10)) + 'T12:00:00')
  if (view.value === 'month') {
    date.setDate(1)
    date.setMonth(date.getMonth() + direction)
  } else date.setDate(date.getDate() + direction * (view.value === 'week' ? 7 : 1))
  day.value = new Date(date.getTime() - date.getTimezoneOffset() * 60000).toISOString().slice(0, 10)
}

function downloadCalendar() {
  const escape = (value) => String(value || '').replace(/\\/g, '\\\\').replace(/\n/g, '\\n').replace(/,/g, '\\,').replace(/;/g, '\\;')
  const timestamp = (date) => new Date(date).toISOString().replace(/[-:]/g, '').replace(/\.\d{3}/, '')
  const lines = ['BEGIN:VCALENDAR', 'VERSION:2.0', 'PRODID:-//UpcycleConnect//Planning//FR']
  for (const entry of entries.value) lines.push('BEGIN:VEVENT', 'UID:' + escape(entry.id) + '@upcycleconnect.local', 'DTSTAMP:' + timestamp(new Date()), 'DTSTART:' + timestamp(entry.start), 'DTEND:' + timestamp(entry.end), 'SUMMARY:' + escape(entry.title), 'LOCATION:' + escape(entry.location), 'BEGIN:VALARM', 'TRIGGER:-P1D', 'ACTION:DISPLAY', 'DESCRIPTION:' + escape(entry.title), 'END:VALARM', 'END:VEVENT')
  lines.push('END:VCALENDAR')
  const url = URL.createObjectURL(new Blob([lines.join('\r\n')], { type: 'text/calendar;charset=utf-8' }))
  const link = document.createElement('a')
  link.href = url
  link.download = 'planning-upcycleconnect.ics'
  link.click()
  setTimeout(() => URL.revokeObjectURL(url), 1000)
}
</script>

<template>
  <div class="employee-section-head"><p>{{ t('employee.planning.intro') }}</p><button v-if="!form" class="e-button" type="button" @click="create">{{ t('employee.planning.new') }}</button></div>
  <p class="e-notice">{{ t('employee.localDraftNotice') }}</p>
  <div class="e-notice"><p>{{ t('employee.planning.browserReminderHint') }}</p><div class="e-actions"><button class="e-button e-secondary" type="button" :disabled="reminderState === 'active'" @click="enableReminders">{{ t('employee.planning.enableReminders') }}</button><span v-if="reminderState !== 'inactive'" role="status">{{ t('employee.planning.reminderStates.' + reminderState) }}</span></div></div>
  <form v-if="form" class="e-panel e-form" @submit.prevent="save">
    <h3>{{ t('employee.planning.formTitle') }}</h3>
    <label>{{ t('employee.training.type') }} <select v-model="form.type"><option v-for="type in types" :key="type" :value="type">{{ t('employee.types.' + type) }}</option></select></label>
    <label>{{ t('employee.training.title') }} <input v-model.trim="form.title" required maxlength="120" /></label>
    <label>{{ t('employee.planning.starts') }} <input v-model="form.start" type="datetime-local" required /></label>
    <label>{{ t('employee.planning.ends') }} <input v-model="form.end" type="datetime-local" required /></label>
    <label class="e-full">{{ t('employee.location') }} <input v-model.trim="form.location" required maxlength="160" /></label>
    <p v-if="error" class="e-full e-error" role="alert">{{ t('employee.' + error) }}</p>
    <div class="e-full e-actions"><button class="e-button" type="submit">{{ t('employee.saveLocalDraft') }}</button><button class="e-button e-secondary" type="button" @click="form = null">{{ t('employee.cancel') }}</button></div>
  </form>
  <template v-else>
    <div class="e-filters">
      <label>{{ t('employee.planning.view') }} <select v-model="view"><option value="day">{{ t('employee.planning.day') }}</option><option value="week">{{ t('employee.planning.week') }}</option><option value="month">{{ t('employee.planning.month') }}</option></select></label>
      <label>{{ t('employee.planning.date') }} <input v-model="day" type="date" required /></label>
      <div class="e-actions"><button class="e-button e-secondary" type="button" @click="move(-1)">{{ t('employee.planning.previous') }}</button><button class="e-button e-secondary" type="button" @click="day = dateAfter(0).slice(0, 10)">{{ t('employee.planning.today') }}</button><button class="e-button e-secondary" type="button" @click="move(1)">{{ t('employee.planning.next') }}</button></div>
    </div>
    <p v-if="day" class="e-period">{{ dateLabel(period.start, locale, false) }} – {{ dateLabel(new Date(period.end.getTime() - 1), locale, false) }}</p>
    <ul v-if="entries.length" class="e-list"><li v-for="item in entries" :key="item.id" class="e-panel"><div class="e-card-head"><h3>{{ item.title }}</h3><span class="e-badge">{{ t('employee.types.' + item.type) }}<template v-if="item.local"> · {{ t('employee.localDraft') }}</template></span></div><p>{{ dateLabel(item.start, locale) }} – {{ dateLabel(item.end, locale) }}</p><p>{{ item.location }}</p><div v-if="item.local && !item.session" class="e-actions"><button class="e-button e-secondary" type="button" @click="form = { ...item }; error = ''">{{ t('employee.edit') }}</button><button class="e-button e-secondary" type="button" @click="remove(item)">{{ t('employee.delete') }}</button></div></li></ul>
    <p v-else class="e-panel">{{ t(remoteUnavailable ? 'employee.planning.emptyDrafts' : 'employee.planning.empty') }}</p>
    <div v-if="entries.length" class="e-actions"><button type="button" class="e-button e-secondary" @click="downloadCalendar">{{ t('employee.planning.export') }}</button><small>{{ t('employee.planning.reminder') }}</small></div>
  </template>
</template>

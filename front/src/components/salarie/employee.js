import { ref, watch } from 'vue'

export function dateAfter(days, hour = '10:00') {
  const date = new Date()
  date.setDate(date.getDate() + days)
  const local = new Date(date.getTime() - date.getTimezoneOffset() * 60000)
  return local.toISOString().slice(0, 10) + 'T' + hour
}

export function dateLabel(value, locale, includeTime = true) {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return '—'
  return new Intl.DateTimeFormat(locale, includeTime
    ? { dateStyle: 'medium', timeStyle: 'short' }
    : { dateStyle: 'medium' }).format(date)
}

export function endOfTraining(training) {
  return new Date(new Date(training.start).getTime() + Number(training.duration) * 60000)
}

export function trainingStatus(training) {
  return training.status
}

export function isConflict(start, end, events, trainings, ignoreId = '') {
  const startTime = new Date(start).getTime()
  const endTime = new Date(end).getTime()
  const intervals = [
    ...events.filter((item) => item.id !== ignoreId).map((item) => [item.start, item.end]),
    ...trainings.filter((item) => item.id !== ignoreId && trainingStatus(item) === 'published')
      .map((item) => [item.start, endOfTraining(item)]),
  ]
  return intervals.some(([from, to]) => startTime < new Date(to).getTime() && endTime > new Date(from).getTime())
}

export function useEmployeeDrafts(scope) {
  const data = ref({ trainings: [], events: [], advice: [] })
  const storageError = ref(false)
  const key = () => 'upcycleconnect:employee:drafts:' + encodeURIComponent(scope())

  watch(scope, (value) => {
    data.value = { trainings: [], events: [], advice: [] }
    storageError.value = false
    if (!value) return
    try {
      const stored = JSON.parse(localStorage.getItem(key()))
      if (stored && ['trainings', 'events', 'advice'].every((item) => Array.isArray(stored[item]))) data.value = stored
    } catch {
      storageError.value = true
    }
  }, { immediate: true })

  function save(name, items) {
    if (!scope()) return false
    const next = { ...data.value, [name]: items }
    try {
      localStorage.setItem(key(), JSON.stringify(next))
      data.value = next
      storageError.value = false
      return true
    } catch {
      storageError.value = true
      return false
    }
  }

  return { data, storageError, save }
}

export function newId() {
  return globalThis.crypto?.randomUUID?.() || Date.now().toString(36) + Math.random().toString(36).slice(2)
}

export function readImage(event, callback, onError) {
  const file = event.target.files?.[0]
  if (!file) return
  if (!['image/jpeg', 'image/png', 'image/webp'].includes(file.type) || file.size > 400 * 1024) {
    onError()
    event.target.value = ''
    return
  }
  const reader = new FileReader()
  reader.onload = () => callback(reader.result)
  reader.onerror = onError
  reader.readAsDataURL(file)
}

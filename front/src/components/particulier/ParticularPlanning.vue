<script setup>
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { demoDeposit, localDate } from './local'

const props = defineProps({ demo: Boolean, bookings: { type: Array, required: true } })
const { t, d } = useI18n()
const modes = ['day', 'week', 'month']
const mode = ref('week')
const date = ref(localDate())

const range = computed(() => {
  const start = new Date((date.value || localDate()) + 'T12:00:00')
  const end = new Date(start)
  if (mode.value === 'week') {
    start.setDate(start.getDate() - (start.getDay() + 6) % 7)
    end.setTime(start.getTime())
    end.setDate(end.getDate() + 6)
  }
  if (mode.value === 'month') {
    start.setDate(1)
    end.setMonth(end.getMonth() + 1, 0)
  }
  return { start, end }
})

const agenda = computed(() => {
  if (!props.demo) return []
  const deposit = demoDeposit()
  return [...props.bookings, {
    id: deposit.id, startsAt: deposit.deadline, location: deposit.location,
  }].sort((a, b) => a.startsAt.localeCompare(b.startsAt))
})

const visible = computed(() => {
  const start = localDate(range.value.start)
  const end = localDate(range.value.end)
  return agenda.value.filter((entry) => {
    const day = entry.startsAt.slice(0, 10)
    return day >= start && day <= end
  })
})

function title(entry) {
  return entry.titleKey
    ? t('particular.catalogue.samples.' + entry.titleKey)
    : t('particular.planning.depositDeadline')
}

function dateLabel(value) {
  return value.length === 10
    ? d(new Date(value + 'T12:00:00'), { dateStyle: 'medium' })
    : d(new Date(value), { dateStyle: 'medium', timeStyle: 'short' })
}

function calendarText(text) {
  return text.replace(/\\/g, '\\\\').replace(/[\r\n]+/g, '\\n').replace(/[;,]/g, '\\$&')
}

function exportCalendar() {
  if (!props.demo || !visible.value.length) return
  const stamp = new Date().toISOString().replace(/[-:]|\.\d{3}/g, '')
  const events = visible.value.flatMap((entry) => {
    const start = entry.startsAt.replace(/[-:]/g, '')
    return [
      'BEGIN:VEVENT',
      'UID:' + entry.id + '-' + start + '@upcycleconnect.local',
      'DTSTAMP:' + stamp,
      (entry.startsAt.length === 10 ? 'DTSTART;VALUE=DATE:' : 'DTSTART:') + start,
      'SUMMARY:' + calendarText(t('particular.demoTitle') + ' — ' + title(entry)),
      'LOCATION:' + calendarText(entry.location),
      'BEGIN:VALARM', 'TRIGGER:-P1D', 'ACTION:DISPLAY',
      'DESCRIPTION:' + calendarText(t('particular.planning.reminder')),
      'END:VALARM', 'END:VEVENT',
    ]
  })

  const calendar = ['BEGIN:VCALENDAR', 'VERSION:2.0', 'PRODID:-//UpcycleConnect//Particulier//FR', ...events, 'END:VCALENDAR', '']
  const url = URL.createObjectURL(new Blob([calendar.join('\r\n')], { type: 'text/calendar;charset=utf-8' }))
  const download = document.createElement('a')
  download.href = url
  download.download = 'upcycleconnect-demo.ics'
  download.click()
  setTimeout(() => URL.revokeObjectURL(url), 1000)
}
</script>

<template>
  <p>{{ t('particular.overview.planning') }}</p>
  <p v-if="!demo" class="p-notice">{{ t('particular.planning.unavailable') }}</p>
  <div class="p-filters">
    <label>{{ t('particular.planning.view') }}
      <select v-model="mode">
        <option v-for="item in modes" :key="item" :value="item">{{ t('particular.planning.' + item) }}</option>
      </select>
    </label>
    <label>{{ t('particular.planning.date') }} <input v-model="date" type="date" /></label>
    <button type="button" class="p-button p-secondary" @click="date = localDate()">{{ t('particular.planning.today') }}</button>
  </div>
  <p class="p-period">
    <strong>{{ d(range.start, { dateStyle: 'long' }) }}</strong>
    <template v-if="mode !== 'day'"> – <strong>{{ d(range.end, { dateStyle: 'long' }) }}</strong></template>
  </p>
  <ul v-if="visible.length" class="p-rows">
    <li v-for="entry in visible" :key="entry.id" class="p-card">
      <h3>{{ title(entry) }}</h3>
      <p>{{ dateLabel(entry.startsAt) }} · {{ entry.location }}</p>
    </li>
  </ul>
  <p v-else-if="demo" class="p-card">{{ t('particular.planning.empty') }}</p>
  <div v-if="visible.length" class="p-actions">
    <button type="button" class="p-button" @click="exportCalendar">{{ t('particular.planning.export') }}</button>
    <p>{{ t('particular.planning.exportHint') }}</p>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { dateAfter, dateLabel, endOfTraining, isConflict, newId, readImage, trainingStatus } from './employee'

const props = defineProps({ items: { type: Array, required: true }, events: { type: Array, required: true }, remoteUnavailable: Boolean, saveItems: { type: Function, required: true } })
const { t, locale, n } = useI18n()
const form = ref(null)
const search = ref('')
const filter = ref('')
const selected = ref('')
const error = ref('')
const types = ['training', 'workshop', 'event']
const levels = ['beginner', 'intermediate', 'expert']
const categories = ['wood', 'metal', 'textile', 'plastic', 'decoration', 'repair', 'tutorials', 'tips']
const visible = computed(() => props.items.filter((item) => {
  const text = `${item.title} ${item.description || ''}`.toLocaleLowerCase()
  return text.includes(search.value.trim().toLocaleLowerCase()) && (!filter.value || trainingStatus(item) === filter.value)
}))
const selectedItem = computed(() => props.items.find((item) => item.id === selected.value))

function create() {
  error.value = ''
  form.value = { id: newId(), local: true, type: 'training', title: '', description: '', category: '', objectives: '', level: 'beginner', start: dateAfter(7), duration: 60, capacity: 8, location: '', price: 0, image: '', status: 'draft', participants: [], waiting: [] }
}

function edit(item, duplicate = false) {
  error.value = ''
  form.value = { ...item, id: duplicate ? newId() : item.id, local: true, level: item.level || 'beginner', category: item.category || '', status: duplicate ? 'draft' : item.status, participants: duplicate ? [] : [...item.participants], waiting: duplicate ? [] : [...item.waiting] }
}

function save() {
  const item = form.value
  const start = new Date(item.start)
  const end = endOfTraining(item)
  if (!item.title.trim() || !item.description.trim() || !item.category || !item.objectives.trim() || !item.location.trim() || Number(item.duration) < 1 || Number(item.capacity) < 1 || Number(item.price) < 0 || Number.isNaN(start.getTime()) || start <= new Date()) {
    error.value = 'formError'
    return
  }
  if (Number(item.capacity) < item.participants.length) {
    error.value = 'capacityError'
    return
  }
  if (isConflict(item.start, end, props.events, props.items, item.id)) {
    error.value = 'conflict'
    return
  }
  const next = { ...item, duration: Number(item.duration), capacity: Number(item.capacity), price: Number(item.price), status: 'draft' }
  if (!props.saveItems([...props.items.filter((entry) => entry.id !== next.id), next])) return
  form.value = null
  error.value = ''
}

function removeDraft(item) {
  if (window.confirm(t('employee.confirmDeleteDraft')) && props.saveItems(props.items.filter((entry) => entry.id !== item.id))) selected.value = ''
}

</script>

<template>
  <div class="employee-section-head">
    <p>{{ t('employee.training.intro') }}</p>
    <button v-if="!form" class="e-button" type="button" @click="create">{{ t('employee.training.new') }}</button>
  </div>
  <p class="e-notice">{{ t('employee.localDraftNotice') }}</p>

  <form v-if="form" class="e-panel e-form" @submit.prevent="save">
    <h3>{{ t('employee.training.formTitle') }}</h3>
    <p v-if="form.status === 'published'" class="e-full e-notice">{{ t('employee.training.revalidation') }}</p>
    <label>{{ t('employee.training.type') }}
      <select v-model="form.type"><option v-for="type in types" :key="type" :value="type">{{ t('employee.types.' + type) }}</option></select>
    </label>
    <label>{{ t('employee.training.title') }} <input v-model.trim="form.title" maxlength="120" required /></label>
    <label class="e-full">{{ t('employee.training.description') }} <textarea v-model.trim="form.description" rows="3" required /></label>
    <label>{{ t('employee.category') }}
      <select v-model="form.category" required><option disabled value="">{{ t('employee.choose') }}</option><option v-for="category in categories" :key="category" :value="category">{{ t('employee.categories.' + category) }}</option></select>
    </label>
    <label>{{ t('employee.training.level') }}
      <select v-model="form.level"><option v-for="level in levels" :key="level" :value="level">{{ t('employee.levels.' + level) }}</option></select>
    </label>
    <label class="e-full">{{ t('employee.training.objectives') }} <textarea v-model.trim="form.objectives" rows="2" required /></label>
    <label>{{ t('employee.training.start') }} <input v-model="form.start" type="datetime-local" required /></label>
    <label>{{ t('employee.training.duration') }} <input v-model.number="form.duration" type="number" min="1" step="1" required /></label>
    <label>{{ t('employee.training.capacity') }} <input v-model.number="form.capacity" type="number" min="1" step="1" required /></label>
    <label>{{ t('employee.training.price') }} <input v-model.number="form.price" type="number" min="0" step="0.01" required /></label>
    <label class="e-full">{{ t('employee.location') }} <input v-model.trim="form.location" maxlength="160" required /></label>
    <label class="e-full">{{ t('employee.image') }} <input type="file" accept="image/jpeg,image/png,image/webp" @change="readImage($event, (image) => form.image = image, () => error = 'imageError')" /><small>{{ t('employee.imageHint') }}</small></label>
    <div v-if="form.image" class="e-full"><img class="e-image" :src="form.image" :alt="t('employee.image')" /><button class="e-text-button" type="button" @click="form.image = ''">{{ t('employee.removeImage') }}</button></div>
    <p v-if="error" class="e-full e-error" role="alert">{{ t('employee.' + error) }}</p>
    <div class="e-actions e-full">
      <button class="e-button" type="submit">{{ t('employee.training.saveDraft') }}</button>
      <button class="e-button e-secondary" type="button" disabled>{{ t('employee.training.submitReal') }}</button>
      <button class="e-button e-secondary" type="button" @click="form = null">{{ t('employee.cancel') }}</button>
    </div>
  </form>

  <template v-else-if="selectedItem">
    <button class="e-text-button" type="button" @click="selected = ''">{{ t('employee.back') }}</button>
    <article class="e-panel">
      <div class="e-card-head"><h3>{{ selectedItem.title }}</h3><span class="e-badge">{{ t('employee.training.status.' + trainingStatus(selectedItem)) }}</span></div>
      <img v-if="selectedItem.image" class="e-image" :src="selectedItem.image" :alt="t('employee.image')" />
      <p>{{ selectedItem.description }}</p>
      <p v-if="selectedItem.status === 'refused'" class="e-notice"><strong>{{ t('employee.training.refusalReason') }} :</strong> {{ selectedItem.rejectionReason || t('employee.training.refusalReasonUnavailable') }}</p>
      <p v-if="selectedItem.category"><strong>{{ t('employee.category') }} :</strong> {{ t('employee.categories.' + selectedItem.category) }}</p>
      <p v-if="selectedItem.level"><strong>{{ t('employee.training.level') }} :</strong> {{ t('employee.levels.' + selectedItem.level) }}</p>
      <p v-if="selectedItem.objectives"><strong>{{ t('employee.training.objectives') }} :</strong> {{ selectedItem.objectives }}</p>
      <p>{{ dateLabel(selectedItem.start, locale) }} · {{ selectedItem.location || '—' }} · {{ selectedItem.duration }} {{ t('employee.minutes') }}</p>
      <p v-if="selectedItem.price != null"><strong>{{ t('employee.training.price') }} :</strong> {{ n(selectedItem.price, { style: 'currency', currency: 'EUR' }) }}</p>
      <p>{{ t('employee.training.registrationsUnavailable') }}</p>
      <div class="e-actions">
        <button v-if="selectedItem.local" class="e-button e-secondary" type="button" @click="edit(selectedItem)">{{ t('employee.edit') }}</button>
        <button class="e-button e-secondary" type="button" @click="edit(selectedItem, true)">{{ t('employee.duplicate') }}</button>
        <button v-if="selectedItem.local" class="e-button e-secondary" type="button" @click="removeDraft(selectedItem)">{{ t('employee.delete') }}</button>
      </div>
    </article>
  </template>

  <template v-else>
    <div class="e-filters">
      <label>{{ t('employee.search') }} <input v-model="search" type="search" /></label>
      <label>{{ t('employee.status') }} <select v-model="filter"><option value="">{{ t('employee.all') }}</option><option v-for="status in ['draft', 'pending', 'validated', 'refused', 'published', 'completed', 'cancelled', 'unknown']" :key="status" :value="status">{{ t('employee.training.status.' + status) }}</option></select></label>
    </div>
    <ul v-if="visible.length" class="e-list"><li v-for="item in visible" :key="item.id" class="e-panel"><div class="e-card-head"><h3>{{ item.title }}</h3><span class="e-badge">{{ t('employee.training.status.' + trainingStatus(item)) }}</span></div><p>{{ t('employee.types.' + item.type) }} · {{ dateLabel(item.start, locale) }} · {{ item.location || '—' }}</p><button class="e-text-button" type="button" @click="selected = item.id">{{ t('employee.details') }}</button></li></ul>
    <p v-else class="e-panel">{{ t(remoteUnavailable ? 'employee.training.emptyDrafts' : 'employee.training.empty') }}</p>
  </template>
</template>

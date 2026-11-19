<script setup>
import { computed, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { dateAfter, dateLabel, newId, readImage } from './employee'

const props = defineProps({ items: { type: Array, required: true }, author: { type: String, default: '' }, saveItems: { type: Function, required: true } })
const { t, locale } = useI18n()
const form = ref(null)
const search = ref('')
const category = ref('')
const status = ref('')
const selected = ref('')
const error = ref('')
const now = ref(Date.now())
const timer = setInterval(() => { now.value = Date.now() }, 30000)
onUnmounted(() => clearInterval(timer))
const categories = ['wood', 'metal', 'textile', 'plastic', 'decoration', 'repair', 'tutorials', 'tips', 'news']
const state = (item) => item.status === 'scheduled' && new Date(item.publishedAt).getTime() <= now.value ? 'published' : item.status
const visible = computed(() => props.items.filter((item) => {
  const text = `${item.title} ${item.content} ${item.keywords}`.toLocaleLowerCase()
  return text.includes(search.value.trim().toLocaleLowerCase()) && (!category.value || item.category === category.value) && (!status.value || state(item) === status.value)
}))
const selectedItem = computed(() => props.items.find((item) => item.id === selected.value))

function create() {
  error.value = ''
  form.value = { id: newId(), local: true, title: '', content: '', category: '', keywords: '', image: '', author: props.author, publishedAt: dateAfter(0), status: 'draft' }
}

function save() {
  const item = form.value
  if (!item.title.trim() || !item.content.trim() || !item.category) {
    error.value = 'formError'
    return
  }
  const next = { ...item, status: 'draft' }
  if (!props.saveItems([...props.items.filter((entry) => entry.id !== item.id), next])) return
  form.value = null
  error.value = ''
}

function remove(item) {
  if (!window.confirm(t('employee.confirmDeleteDraft'))) return
  if (props.saveItems(props.items.filter((entry) => entry.id !== item.id))) selected.value = ''
}
</script>

<template>
  <div class="employee-section-head"><p>{{ t('employee.advice.intro') }}</p><button v-if="!form" class="e-button" type="button" @click="create">{{ t('employee.advice.new') }}</button></div>
  <p class="e-notice">{{ t('employee.localDraftNotice') }}</p>
  <form v-if="form" class="e-panel e-form" @submit.prevent="save">
    <h3>{{ t('employee.advice.formTitle') }}</h3>
    <label>{{ t('employee.training.title') }} <input v-model.trim="form.title" maxlength="120" required /></label>
    <label>{{ t('employee.category') }} <select v-model="form.category" required><option value="" disabled>{{ t('employee.choose') }}</option><option v-for="item in categories" :key="item" :value="item">{{ t('employee.categories.' + item) }}</option></select></label>
    <label class="e-full">{{ t('employee.advice.content') }} <textarea v-model.trim="form.content" rows="6" required /></label>
    <label>{{ t('employee.advice.keywords') }} <input v-model.trim="form.keywords" maxlength="200" /></label>
    <label>{{ t('employee.advice.scheduledFor') }} <input v-model="form.publishedAt" type="datetime-local" /></label>
    <label class="e-full">{{ t('employee.image') }} <input type="file" accept="image/jpeg,image/png,image/webp" @change="readImage($event, (image) => form.image = image, () => error = 'imageError')" /><small>{{ t('employee.imageHint') }}</small></label>
    <div v-if="form.image" class="e-full"><img class="e-image" :src="form.image" :alt="t('employee.image')" /><button class="e-text-button" type="button" @click="form.image = ''">{{ t('employee.removeImage') }}</button></div>
    <p v-if="error" class="e-full e-error" role="alert">{{ t('employee.' + error) }}</p>
    <div class="e-full e-actions"><button class="e-button" type="submit">{{ t('employee.advice.saveDraft') }}</button><button class="e-button e-secondary" type="button" disabled>{{ t('employee.advice.publishReal') }}</button><button class="e-button e-secondary" type="button" disabled>{{ t('employee.advice.scheduleReal') }}</button><button class="e-button e-secondary" type="button" @click="form = null">{{ t('employee.cancel') }}</button></div>
  </form>
  <template v-else-if="selectedItem">
    <button class="e-text-button" type="button" @click="selected = ''">{{ t('employee.back') }}</button>
    <article class="e-panel"><div class="e-card-head"><h3>{{ selectedItem.title }}</h3><span class="e-badge">{{ t('employee.advice.status.' + state(selectedItem)) }}</span></div><p>{{ t('employee.categories.' + selectedItem.category) }} · {{ selectedItem.author }} · {{ dateLabel(selectedItem.publishedAt, locale) }}</p><img v-if="selectedItem.image" class="e-image" :src="selectedItem.image" :alt="t('employee.image')" /><p class="e-content">{{ selectedItem.content }}</p><p v-if="selectedItem.keywords">{{ t('employee.advice.keywords') }} : {{ selectedItem.keywords }}</p><div v-if="selectedItem.local" class="e-actions"><button class="e-button e-secondary" type="button" @click="form = { ...selectedItem }; error = ''">{{ t('employee.edit') }}</button><button class="e-button e-secondary" type="button" @click="remove(selectedItem)">{{ t('employee.delete') }}</button></div></article>
  </template>
  <template v-else>
    <div class="e-filters"><label>{{ t('employee.search') }} <input v-model="search" type="search" /></label><label>{{ t('employee.category') }} <select v-model="category"><option value="">{{ t('employee.all') }}</option><option v-for="item in categories" :key="item" :value="item">{{ t('employee.categories.' + item) }}</option></select></label><label>{{ t('employee.status') }} <select v-model="status"><option value="">{{ t('employee.all') }}</option><option v-for="item in ['draft', 'scheduled', 'published', 'archived']" :key="item" :value="item">{{ t('employee.advice.status.' + item) }}</option></select></label></div>
    <ul v-if="visible.length" class="e-list"><li v-for="item in visible" :key="item.id" class="e-panel"><div class="e-card-head"><h3>{{ item.title }}</h3><span class="e-badge">{{ t('employee.advice.status.' + state(item)) }}</span></div><p>{{ t('employee.categories.' + item.category) }} · {{ item.author }} · {{ dateLabel(item.publishedAt, locale) }}</p><button class="e-text-button" type="button" @click="selected = item.id">{{ t('employee.details') }}</button></li></ul>
    <p v-else class="e-panel">{{ t('employee.advice.emptyDrafts') }}</p>
  </template>
</template>

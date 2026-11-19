<script setup>
import { computed, nextTick, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { demoDeposit, useLocalValue } from './local'

const props = defineProps({
  kind: { type: String, required: true },
  scope: { type: String, required: true },
  demo: Boolean,
})
const { t, n, d } = useI18n()
const listing = computed(() => props.kind === 'listings')
const { value: records, save } = useLocalValue(() => props.scope, () => props.kind, [])
const visibleRecords = computed(() => props.demo && !listing.value ? [...records.value, demoDeposit()] : records.value)
const categories = ['wood', 'metal', 'textile', 'plastic', 'other']
const conditions = ['good', 'used', 'repair']
const form = ref(null)
const firstField = ref(null)
const reading = ref(false)
const feedback = ref('')
const error = ref('')

async function edit(record) {
  error.value = ''
  feedback.value = ''
  form.value = record
    ? { ...record, photos: [...record.photos] }
    : { id: crypto.randomUUID(), title: '', description: '', category: '', condition: '',
        location: '', type: 'donation', price: '', dimensions: '', weight: '', photos: [], status: 'draft' }
  await nextTick()
  firstField.value?.focus()
}

function store(status) {
  error.value = ''
  const draft = form.value
  if (!draft.description.trim() || !draft.location.trim()
    || (listing.value && !draft.title.trim()) || (!listing.value && !draft.dimensions.trim())) {
    error.value = 'required'
    return
  }
  if (status === 'pending' && !draft.photos.length) {
    error.value = 'photoRequired'
    return
  }

  const updated = { ...draft, status, price: draft.type === 'sale' ? draft.price : '' }
  const next = records.value.some((item) => item.id === updated.id)
    ? records.value.map((item) => item.id === updated.id ? updated : item)
    : [...records.value, updated]

  if (!save(next)) {
    error.value = 'storageError'
    return
  }
  form.value = null
  feedback.value = status === 'draft' ? 'draftSaved' : 'demoSubmitted'
}

function remove(record) {
  if (!window.confirm(t('particular.requests.confirmRemove'))) return
  error.value = ''
  if (save(records.value.filter((item) => item.id !== record.id))) {
    if (form.value?.id === record.id) form.value = null
    feedback.value = 'removed'
  } else {
    error.value = 'storageError'
  }
}

function readPhoto(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = () => resolve(reader.result)
    reader.onerror = reject
    reader.readAsDataURL(file)
  })
}

async function addPhotos(event) {
  const files = Array.from(event.target.files)
  event.target.value = ''
  error.value = ''
  if (form.value.photos.length + files.length > 3
    || files.some((file) => !['image/jpeg', 'image/png', 'image/webp'].includes(file.type) || file.size > 400 * 1024)) {
    error.value = 'photoLimit'
    return
  }
  reading.value = true
  try {
    form.value.photos.push(...await Promise.all(files.map(readPhoto)))
  } catch {
    error.value = 'photoError'
  } finally {
    reading.value = false
  }
}
</script>

<template>
  <p>{{ t('particular.overview.' + kind) }}</p>
  <p v-if="!demo" class="p-notice">{{ t('particular.requests.draftsOnly') }}</p>
  <div v-if="!form" class="p-actions">
    <button type="button" class="p-button" @click="edit()">
      {{ t('particular.requests.' + (listing ? 'newListing' : 'newDeposit')) }}
    </button>
  </div>

  <form v-else class="p-card p-form" @submit.prevent="store($event.submitter?.value === 'pending' ? 'pending' : 'draft')">
    <h3>{{ t('particular.requests.' + (listing ? 'listingForm' : 'depositForm')) }}</h3>
    <label v-if="listing">{{ t('particular.requests.title') }}
      <input ref="firstField" v-model.trim="form.title" required maxlength="100" />
    </label>
    <label class="p-wide">{{ t('particular.requests.description') }}
      <textarea v-if="listing" v-model.trim="form.description" rows="3" required maxlength="1000"></textarea>
      <textarea v-else ref="firstField" v-model.trim="form.description"
        rows="3" required maxlength="1000"></textarea>
    </label>
    <label>{{ t('particular.requests.category') }}
      <select v-model="form.category" required>
        <option disabled value="">{{ t('particular.requests.choose') }}</option>
        <option v-for="category in categories" :key="category" :value="category">
          {{ t('particular.categories.' + category) }}
        </option>
      </select>
    </label>
    <label>{{ t('particular.requests.location') }}
      <input v-model.trim="form.location" required maxlength="120" />
    </label>

    <template v-if="listing">
      <label>{{ t('particular.requests.condition') }}
        <select v-model="form.condition" required>
          <option disabled value="">{{ t('particular.requests.choose') }}</option>
          <option v-for="condition in conditions" :key="condition" :value="condition">
            {{ t('particular.requests.' + condition) }}
          </option>
        </select>
      </label>
      <label>{{ t('particular.requests.type') }}
        <select v-model="form.type">
          <option value="donation">{{ t('particular.requests.donation') }}</option>
          <option value="sale">{{ t('particular.requests.sale') }}</option>
        </select>
      </label>
      <label v-if="form.type === 'sale'">{{ t('particular.requests.price') }}
        <input v-model.number="form.price" type="number" min="0.01" max="99999" step="0.01" required />
      </label>
    </template>
    <template v-else>
      <label>{{ t('particular.requests.dimensions') }}
        <input v-model.trim="form.dimensions" required maxlength="80" />
      </label>
      <label>{{ t('particular.requests.weight') }}
        <input v-model.number="form.weight" type="number" min="0.01" step="0.01" required />
      </label>
    </template>

    <label class="p-wide">{{ t('particular.requests.photos') }}
      <input type="file" accept="image/jpeg,image/png,image/webp" multiple
        aria-describedby="particular-photo-limit" :disabled="reading" @change="addPhotos" />
      <small id="particular-photo-limit">{{ t('particular.requests.photoLimit') }}</small>
    </label>
    <div v-if="form.photos.length" class="p-actions p-wide p-photos">
      <div v-for="(photo, index) in form.photos" :key="index">
        <img :src="photo" :alt="t('particular.requests.photo', { number: index + 1 })" />
        <button type="button" class="p-button p-secondary" :disabled="reading" @click="form.photos.splice(index, 1)">
          {{ t('particular.requests.removePhoto', { number: index + 1 }) }}
        </button>
      </div>
    </div>
    <div class="p-actions p-wide">
      <button v-if="form.status === 'draft'" type="submit" value="draft" class="p-button" :disabled="reading">
        {{ t('particular.requests.saveDraft') }}
      </button>
      <button v-if="demo" type="submit" value="pending" class="p-button p-secondary" :disabled="reading">
        {{ t(form.status === 'pending' ? 'particular.requests.updateDemo' : 'particular.requests.submitDemo') }}
      </button>
      <button type="button" class="p-button p-secondary" :disabled="reading" @click="form = null; error = ''">
        {{ t('admin.cancel') }}
      </button>
    </div>
  </form>

  <p v-if="error" class="p-notice" role="alert">
    {{ t(error === 'storageError' ? 'particular.storageError' : 'particular.requests.' + error) }}
  </p>
  <p v-if="feedback" class="p-notice" role="status">{{ t('particular.requests.' + feedback) }}</p>

  <ul v-if="visibleRecords.length" class="p-rows">
    <li v-for="record in visibleRecords" :key="record.id" class="p-card">
      <div class="p-card-heading">
        <h3>{{ record.title || t('particular.categories.' + record.category) }}</h3>
        <span class="p-chip">{{ t('particular.requests.states.' + record.status) }}</span>
      </div>
      <p>{{ record.location }}</p>
      <p v-if="record.description">{{ record.description }}</p>
      <p v-if="listing">
        {{ record.type === 'sale' ? n(record.price, { style: 'currency', currency: 'EUR' }) : t('particular.requests.donation') }}
      </p>
      <dl v-if="record.status === 'approved' && !listing" class="p-grid p-profile">
        <div><dt>{{ t('particular.requests.code') }}</dt><dd>{{ record.code }}</dd></div>
        <div><dt>{{ t('particular.requests.barcode') }}</dt><dd>{{ record.barcode }}</dd></div>
        <div><dt>{{ t('particular.requests.deadline') }}</dt><dd>{{ d(new Date(record.deadline + 'T12:00:00'), { dateStyle: 'medium' }) }}</dd></div>
      </dl>
      <div v-if="['draft', 'pending'].includes(record.status)" class="p-actions">
        <button type="button" class="p-button p-secondary" :disabled="reading" @click="edit(record)">{{ t('admin.edit') }}</button>
        <button type="button" class="p-button p-secondary" :disabled="reading" @click="remove(record)">{{ t('admin.delete') }}</button>
      </div>
    </li>
  </ul>
  <p v-else-if="!form" class="p-card">{{ t('particular.requests.' + (listing ? 'noListings' : 'noDeposits')) }}</p>
</template>

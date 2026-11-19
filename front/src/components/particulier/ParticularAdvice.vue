<script setup>
import { computed, ref, watch } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useLocalValue } from './local'

const props = defineProps({ scope: { type: String, required: true }, demo: Boolean })
const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const { value: favourites, save } = useLocalValue(() => props.scope, 'favourites', [])
const categories = ['wood', 'metal', 'textile', 'plastic', 'tutorials', 'ideas']
const search = ref('')
const category = ref('')
const onlySaved = ref(false)
const sharedLink = ref('')
const copied = ref(false)
const error = ref(false)
const selected = computed(() => categories.find((item) => item === route.query.advice))

function linkTo(item) {
  return { name: 'particular-space', query: { section: 'advice', demo: '1', advice: item } }
}

function article(item, field) {
  return t('particular.advice.articles.' + item + '.' + field)
}

const visible = computed(() => {
  const query = search.value.trim().toLocaleLowerCase()
  return categories.filter((item) => (
    (!category.value || item === category.value)
    && (!onlySaved.value || favourites.value.includes(item))
    && (article(item, 'title') + ' ' + article(item, 'text')).toLocaleLowerCase().includes(query)
  ))
})

function toggleFavourite(item) {
  const next = favourites.value.includes(item)
    ? favourites.value.filter((value) => value !== item)
    : [...favourites.value, item]
  error.value = !save(next)
}

async function share(item) {
  const url = new URL(router.resolve(linkTo(item)).href, window.location.origin).href
  sharedLink.value = url
  copied.value = false
  try {
    await navigator.clipboard.writeText(url)
    if (selected.value === item && sharedLink.value === url) copied.value = true
  } catch {
    // The displayed link can also be copied manually.
  }
}

watch(selected, () => {
  sharedLink.value = ''
  copied.value = false
})
</script>

<template>
  <p>{{ t('particular.overview.advice') }}</p>
  <p v-if="!demo" class="p-notice">{{ t('particular.advice.unavailable') }}</p>

  <template v-else>
    <p v-if="error" role="alert" class="p-notice">{{ t('particular.storageError') }}</p>

    <template v-if="selected">
      <RouterLink :to="linkTo(undefined)" class="p-back">{{ t('particular.advice.back') }}</RouterLink>
      <article class="p-card">
        <span class="p-chip">{{ t('particular.categories.' + selected) }}</span>
        <h3>{{ article(selected, 'title') }}</h3>
        <p>{{ article(selected, 'text') }}</p>
        <div class="p-actions">
          <button type="button" class="p-button p-secondary" :aria-pressed="favourites.includes(selected)" @click="toggleFavourite(selected)">
            {{ t(favourites.includes(selected) ? 'particular.advice.removeFavourite' : 'particular.advice.addFavourite') }}
          </button>
          <button type="button" class="p-button p-secondary" @click="share(selected)">
            {{ t('particular.advice.share') }}
          </button>
        </div>
        <label v-if="sharedLink" class="p-share">{{ t('particular.advice.shareLink') }}
          <input :value="sharedLink" readonly @focus="$event.target.select()" />
        </label>
        <p v-if="copied" role="status">{{ t('particular.advice.copied') }}</p>
      </article>
    </template>

    <template v-else>
      <div class="p-filters">
        <label>{{ t('particular.search') }} <input v-model="search" type="search" /></label>
        <label>{{ t('particular.requests.category') }}
          <select v-model="category">
            <option value="">{{ t('particular.all') }}</option>
            <option v-for="item in categories" :key="item" :value="item">{{ t('particular.categories.' + item) }}</option>
          </select>
        </label>
        <label class="p-checkbox"><input v-model="onlySaved" type="checkbox" /> {{ t('particular.advice.onlySaved') }}</label>
      </div>

      <ul v-if="visible.length" class="p-grid p-rows">
        <li v-for="item in visible" :key="item" class="p-card">
          <span class="p-chip">{{ t('particular.categories.' + item) }}</span>
          <h3><RouterLink :to="linkTo(item)">{{ article(item, 'title') }}</RouterLink></h3>
          <p v-if="favourites.includes(item)">{{ t('particular.advice.saved') }}</p>
        </li>
      </ul>
      <p v-else class="p-card">{{ t('particular.advice.empty') }}</p>
    </template>
  </template>
</template>

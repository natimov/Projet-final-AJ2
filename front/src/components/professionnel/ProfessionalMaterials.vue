<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const search = ref('')
const category = ref('')
const categories = ['wood', 'metal', 'textile', 'plastic', 'other']
</script>

<template>
  <div class="pro-project-form">
    <label>{{ t('professional.materials.search') }}<input v-model="search" type="search" /></label>
    <label
      >{{ t('professional.projects.categoryLabel') }}
      <select v-model="category">
        <option value="">{{ t('professional.materials.all') }}</option>
        <option v-for="item in categories" :key="item" :value="item">
          {{ t('professional.materials.categories.' + item) }}
        </option>
      </select>
    </label>
  </div>
  <p class="pro-empty-state">{{ t('professional.materials.unavailable') }}</p>
</template>

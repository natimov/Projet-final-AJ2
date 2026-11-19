<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const tab = ref('reports')
</script>

<template>
  <p>{{ t('employee.moderation.intro') }}</p>
  <div class="e-actions" :aria-label="t('employee.moderation.view')">
    <button type="button" class="e-button e-secondary" :aria-pressed="tab === 'reports'" @click="tab = 'reports'">{{ t('employee.moderation.reports') }}</button>
    <button type="button" class="e-button e-secondary" :aria-pressed="tab === 'discussions'" @click="tab = 'discussions'">{{ t('employee.moderation.discussions') }}</button>
  </div>
  <p class="e-notice" role="status">{{ t('employee.api.unavailableSection') }}</p>
</template>

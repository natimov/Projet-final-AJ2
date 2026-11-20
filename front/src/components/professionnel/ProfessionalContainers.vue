<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const code = ref('')
</script>
<template>
  <div class="pro-panel pro-project-form">
    <label>{{ t('professional.containers.code') }}<input v-model="code" /></label>
    <div class="pro-form-full">
      <button class="pro-button" type="button" disabled>
        {{ t('professional.containers.validate') }}
      </button>
    </div>
    <p class="pro-form-full">{{ t('professional.containers.unavailable') }}</p>
  </div>
</template>

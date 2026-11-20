<script setup>
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
</script>
<template>
  <div class="pro-overview-grid">
    <section v-for="item in ['plan', 'contracts', 'invoices']" :key="item" class="pro-panel">
      <h3>{{ t('professional.subscription.' + item) }}</h3>
      <p class="pro-panel-message">{{ t('professional.subscription.unavailable') }}</p>
    </section>
  </div>
</template>

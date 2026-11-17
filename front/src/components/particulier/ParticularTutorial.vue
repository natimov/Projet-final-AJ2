<script setup>
import { computed, nextTick, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({ firstVisit: Boolean })
const emit = defineEmits(['finished'])
const { t } = useI18n()
const dialog = ref(null)
const stepHeading = ref(null)
const step = ref(0)
const steps = ['overview', 'listings', 'deposits', 'advice', 'catalogue', 'planning', 'score']
const current = computed(() => steps[step.value])

function open() {
  step.value = 0
  if (dialog.value && !dialog.value.open) dialog.value.showModal()
}

function finish() {
  dialog.value?.close()
  emit('finished')
}

async function move(amount) {
  if (step.value + amount === steps.length) {
    finish()
    return
  }
  step.value += amount
  await nextTick()
  stepHeading.value?.focus()
}

onMounted(() => {
  if (props.firstVisit) open()
})
defineExpose({ open })
</script>

<template>
  <dialog ref="dialog" class="p-tutorial" aria-labelledby="particular-guide-title" @cancel.prevent="finish">
    <p>{{ t('particular.tutorial.progress', { current: step + 1, total: steps.length }) }}</p>
    <h2 id="particular-guide-title" ref="stepHeading" tabindex="-1">{{ t('particular.sections.' + current) }}</h2>
    <p>{{ t('particular.tutorial.steps.' + current) }}</p>
    <div class="p-actions">
      <button type="button" class="p-button p-secondary" autofocus @click="finish">{{ t('particular.tutorial.skip') }}</button>
      <button v-if="step" type="button" class="p-button p-secondary" @click="move(-1)">{{ t('particular.tutorial.previous') }}</button>
      <button type="button" class="p-button" @click="move(1)">
        {{ t(step === steps.length - 1 ? 'particular.tutorial.finish' : 'particular.tutorial.next') }}
      </button>
    </div>
  </dialog>
</template>

import { ref, toValue, watch } from 'vue'

export function useLocalValue(scope, name, fallback) {
  const value = ref(fallback)
  const key = () => 'upcycleconnect:particulier:' + encodeURIComponent(toValue(scope)) + ':' + toValue(name)

  watch(() => toValue(scope), () => {
    try {
      value.value = JSON.parse(localStorage.getItem(key())) ?? fallback
    } catch {
      value.value = fallback
    }
  }, { immediate: true, flush: 'sync' })

  function save(next) {
    try {
      localStorage.setItem(key(), JSON.stringify(next))
      value.value = next
      return true
    } catch {
      return false
    }
  }

  return { value, save }
}

export function localDate(date = new Date()) {
  return [date.getFullYear(), date.getMonth() + 1, date.getDate()]
    .map((part) => String(part).padStart(2, '0')).join('-')
}

export function dateInDays(days) {
  const date = new Date()
  date.setDate(date.getDate() + days)
  return localDate(date)
}

export function demoDeposit() {
  return {
    id: 'demo-deposit', category: 'wood', location: 'Paris 11e', status: 'approved',
    photos: [], deadline: dateInDays(5), code: 'DEMO-4821', barcode: 'DEMO-004821',
  }
}

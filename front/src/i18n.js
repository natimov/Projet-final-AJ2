import { createI18n } from 'vue-i18n'

const localeFiles = import.meta.glob('./locales/*.json', {
  eager: true,
  import: 'default',
})

const messages = {}

for (const path in localeFiles) {
  const locale = path.split('/').pop().replace('.json', '')
  messages[locale] = localeFiles[path]
}

const savedLocale = localStorage.getItem('language')
const initialLocale = messages[savedLocale] ? savedLocale : 'fr'

const i18n = createI18n({
  legacy: false,
  locale: initialLocale,
  fallbackLocale: 'fr',
  messages,
})

export default i18n
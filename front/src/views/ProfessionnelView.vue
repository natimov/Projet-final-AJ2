<script setup>
import { computed, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAuthStore } from '../stores/auth'
import ProfessionalMaterials from '../components/professionnel/ProfessionalMaterials.vue'
import ProfessionalContainers from '../components/professionnel/ProfessionalContainers.vue'
import ProfessionalProjects from '../components/professionnel/ProfessionalProjects.vue'
import ProfessionalSubscription from '../components/professionnel/ProfessionalSubscription.vue'
import '../assets/professionnel.css'

const route = useRoute()
const auth = useAuthStore()
const { t } = useI18n()
const sections = ['overview', 'materials', 'containers', 'projects', 'subscription']
const section = computed(() =>
  sections.includes(route.query.section) ? route.query.section : 'overview',
)
const displayName = computed(() => auth.user?.prenom || auth.user?.nom || t('professional.member'))
const projects = ref([])

const metrics = computed(() => [
  { key: 'projects', value: projects.value.length },
  {
    key: 'completed',
    value: projects.value.filter((project) => project.statut === 'termine').length,
  },
])

function sectionLink(item) {
  return { name: 'professional-space', query: { section: item === 'overview' ? undefined : item } }
}
</script>

<template>
  <main class="professional-page">
    <div class="professional-shell">
      <header class="pro-heading">
        <div>
          <p class="pro-eyebrow">{{ t('professional.kicker') }}</p>
          <h1>{{ t('professional.greeting', { name: displayName }) }}</h1>
          <p class="pro-heading-copy">{{ t('professional.intro') }}</p>
        </div>
        <span class="pro-account-chip"
          ><span aria-hidden="true" class="pro-account-dot"></span
          >{{ t('professional.accountType') }}</span
        >
      </header>

      <div class="pro-layout">
        <nav class="pro-nav" :aria-label="t('professional.navigation')">
          <RouterLink
            v-for="(item, index) in sections"
            :key="item"
            :to="sectionLink(item)"
            :aria-current="section === item ? 'page' : null"
            :class="{ active: section === item }"
          >
            <span class="pro-nav-index" aria-hidden="true">{{ index + 1 }}</span>
            {{ t('professional.sections.' + item) }}
          </RouterLink>
        </nav>

        <section class="pro-content" aria-labelledby="pro-section-title">
          <div class="pro-section-heading">
            <div>
              <p class="pro-eyebrow">{{ t('professional.sectionLabel') }}</p>
              <h2 id="pro-section-title">{{ t('professional.sections.' + section) }}</h2>
              <p>{{ t('professional.descriptions.' + section) }}</p>
            </div>
            <RouterLink
              v-if="section === 'overview'"
              class="pro-button"
              :to="sectionLink('projects')"
            >
              {{ t('professional.actions.viewProjects') }}
            </RouterLink>
          </div>

          <template v-if="section === 'overview'">
            <dl class="pro-metrics">
              <div v-for="metric in metrics" :key="metric.key" class="pro-metric">
                <dt>{{ t('professional.overview.metrics.' + metric.key) }}</dt>
                <dd>{{ metric.value }}</dd>
              </div>
            </dl>

            <div class="pro-overview-grid">
              <section class="pro-panel">
                <div class="pro-panel-heading">
                  <div>
                    <p class="pro-eyebrow">{{ t('professional.overview.projectsEyebrow') }}</p>
                    <h3>{{ t('professional.overview.projectsTitle') }}</h3>
                  </div>
                  <RouterLink class="pro-text-link" :to="sectionLink('projects')">{{
                    t('professional.actions.viewAll')
                  }}</RouterLink>
                </div>
                <ul v-if="projects.length" class="pro-compact-list">
                  <li v-for="project in projects.slice(0, 3)" :key="project.ID">
                    <span class="pro-material-mark" aria-hidden="true">{{
                      t('professional.materials.categories.' + project.categorie).slice(0, 1)
                    }}</span>
                    <span class="pro-compact-copy">
                      <strong>{{ project.titre }}</strong>
                      <small>{{ t('professional.projects.status.' + project.statut) }}</small>
                    </span>
                  </li>
                </ul>
                <p v-else class="pro-panel-message" role="status">
                  {{ t('professional.states.empty') }}
                </p>
              </section>

              <section class="pro-panel">
                <div class="pro-panel-heading">
                  <div>
                    <p class="pro-eyebrow">{{ t('professional.overview.marketEyebrow') }}</p>
                    <h3>{{ t('professional.sections.materials') }}</h3>
                  </div>
                  <RouterLink class="pro-text-link" :to="sectionLink('materials')">{{
                    t('professional.actions.viewAll')
                  }}</RouterLink>
                </div>
                <p class="pro-panel-message">{{ t('professional.materials.unavailable') }}</p>
              </section>
            </div>
          </template>

          <ProfessionalMaterials v-else-if="section === 'materials'" />
          <ProfessionalContainers v-else-if="section === 'containers'" />
          <ProfessionalProjects
            v-else-if="section === 'projects'"
            :projects="projects"
            @changed="projects = $event"
          />
          <ProfessionalSubscription v-else-if="section === 'subscription'" />
        </section>
      </div>
    </div>
  </main>
</template>

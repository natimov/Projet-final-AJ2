<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

defineProps({ projects: { type: Array, required: true } })
const emit = defineEmits(['changed'])
const { t } = useI18n()
const categories = ['wood', 'metal', 'textile', 'plastic', 'other']
const statuses = ['en_preparation', 'en_cours', 'termine']
const showForm = ref(false)
const editingId = ref(null)
const draft = ref({})

function edit(project) {
  editingId.value = project?.ID ?? null
  draft.value = project
    ? { ...project }
    : {
        titre: '',
        description: '',
        categorie: '',
        statut: 'en_preparation',
        impact: '',
        photoAvant: '',
        photoApres: '',
        featured: false,
      }
  showForm.value = true
}

function photo(event, field) {
  const file = event.target.files?.[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = () => {
    draft.value[field] = reader.result
  }
  reader.readAsDataURL(file)
}

function save(projects) {
  const item = { ...draft.value, ID: editingId.value || Date.now() }
  emit(
    'changed',
    editingId.value
      ? projects.map((project) => (project.ID === editingId.value ? item : project))
      : [item, ...projects],
  )
  showForm.value = false
}

function remove(projects, id) {
  if (
    window.confirm(
      t('professional.projects.confirmDelete', { name: projects.find((p) => p.ID === id).titre }),
    )
  ) {
    emit(
      'changed',
      projects.filter((project) => project.ID !== id),
    )
  }
}
</script>

<template>
  <div class="pro-section-actions">
    <button v-if="!showForm" class="pro-button" type="button" @click="edit()">
      {{ t('professional.projects.new') }}
    </button>
  </div>
  <form v-if="showForm" class="pro-panel pro-project-form" @submit.prevent="save(projects)">
    <h3>
      {{ t(editingId ? 'professional.projects.editTitle' : 'professional.projects.formTitle') }}
    </h3>
    <label
      >{{ t('professional.projects.titleLabel')
      }}<input v-model.trim="draft.titre" maxlength="100" required
    /></label>
    <label
      >{{ t('professional.projects.categoryLabel') }}
      <select v-model="draft.categorie" required>
        <option value="" disabled>{{ t('professional.projects.chooseCategory') }}</option>
        <option v-for="item in categories" :key="item" :value="item">
          {{ t('professional.materials.categories.' + item) }}
        </option>
      </select>
    </label>
    <label
      >{{ t('professional.projects.statusLabel') }}
      <select v-model="draft.statut">
        <option v-for="status in statuses" :key="status" :value="status">
          {{ t('professional.projects.status.' + status) }}
        </option>
      </select>
    </label>
    <label
      >{{ t('professional.projects.impactLabel')
      }}<input v-model.number="draft.impact" type="number" min="0" step="0.1"
    /></label>
    <label
      >{{ t('professional.projects.beforePhoto')
      }}<input type="file" accept="image/*" @change="photo($event, 'photoAvant')"
    /></label>
    <label
      >{{ t('professional.projects.afterPhoto')
      }}<input type="file" accept="image/*" @change="photo($event, 'photoApres')"
    /></label>
    <label class="pro-form-full"
      >{{ t('professional.projects.descriptionLabel')
      }}<textarea v-model.trim="draft.description" rows="4" maxlength="800" required />
    </label>
    <div class="pro-form-full pro-actions">
      <button class="pro-button" type="submit">{{ t('professional.projects.save') }}</button
      ><button class="pro-button pro-button-secondary" type="button" @click="showForm = false">
        {{ t('professional.projects.cancel') }}
      </button>
    </div>
  </form>
  <ul v-if="projects.length" class="pro-project-grid">
    <li v-for="project in projects" :key="project.ID" class="pro-project-card">
      <div class="pro-project-art" :class="'category-' + project.categorie" aria-hidden="true">
        <span>{{ t('professional.materials.categories.' + project.categorie).slice(0, 1) }}</span>
      </div>
      <div class="pro-project-card-body">
        <div class="pro-project-meta">
          <span class="pro-status">{{
            t('professional.materials.categories.' + project.categorie)
          }}</span
          ><span class="pro-status">{{ t('professional.projects.status.' + project.statut) }}</span>
        </div>
        <h3>{{ project.titre }}</h3>
        <p>{{ project.description }}</p>
        <p v-if="project.impact !== ''">
          {{ t('professional.projects.impactValue', { weight: project.impact }) }}
        </p>
        <img
          v-if="project.photoAvant"
          :src="project.photoAvant"
          :alt="t('professional.projects.beforePhoto')"
          width="180"
        />
        <img
          v-if="project.photoApres"
          :src="project.photoApres"
          :alt="t('professional.projects.afterPhoto')"
          width="180"
        />
        <div class="pro-actions">
          <button class="pro-button pro-button-secondary" type="button" disabled>
            {{ t('professional.projects.feature') }}
          </button>
          <button class="pro-button pro-button-secondary" type="button" @click="edit(project)">
            {{ t('professional.projects.edit') }}
          </button>
          <button
            class="pro-button pro-button-secondary"
            type="button"
            @click="remove(projects, project.ID)"
          >
            {{ t('professional.projects.delete') }}
          </button>
        </div>
      </div>
    </li>
  </ul>
  <p v-else-if="!showForm" class="pro-empty-state">{{ t('professional.states.empty') }}</p>
</template>

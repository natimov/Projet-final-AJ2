<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'

const users = ref([])
const errorMessage = ref('')

const authStore = useAuthStore()

async function fetchUsers() {
  errorMessage.value = ''

  try {
    const response = await fetch('http://localhost:8080/users', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })

    if (!response.ok) {
      throw new Error('Impossible de récupérer les utilisateurs')
    }

    users.value = await response.json()
  } catch (error) {
    errorMessage.value = error.message
  }
}
async function deleteUser(id) {
  if (!confirm('Supprimer cet utilisateur ?')) {
    return
  }

  try {
    const response = await fetch(`http://localhost:8080/users/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${authStore.token}` },
    })

    if (!response.ok) {
      throw new Error('Impossible de supprimer cet utilisateur')
    }

    users.value = users.value.filter((user) => user.ID !== id)
  } catch (error) {
    errorMessage.value = error.message
  }
}
const editingUserId = ref(null)
const editForm = ref({ nom: '', prenom: '', email: '', role: '' })

function startEdit(user) {
  editingUserId.value = user.ID
  editForm.value = { nom: user.nom, prenom: user.prenom, email: user.email, role: user.role }
}

function cancelEdit() {
  editingUserId.value = null
}

async function saveEdit(id) {
  try {
    const response = await fetch(`http://localhost:8080/users/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify(editForm.value),
    })

    if (!response.ok) {
      throw new Error('Impossible de modifier cet utilisateur')
    }

    const updatedUser = await response.json()
    const index = users.value.findIndex((user) => user.ID === id)
    users.value[index] = updatedUser
    editingUserId.value = null
  } catch (error) {
    errorMessage.value = error.message
  }
}

onMounted(fetchUsers)
</script>

<template>
  <div>
    <h1>Utilisateurs</h1>

    <p v-if="errorMessage">{{ errorMessage }}</p>

    <table v-else>
      <thead>
        <tr>
          <th>Nom</th>
          <th>Prénom</th>
          <th>Email</th>
          <th>Rôle</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.ID">
          <template v-if="editingUserId === user.ID">
            <td><input v-model="editForm.nom" /></td>
            <td><input v-model="editForm.prenom" /></td>
            <td><input v-model="editForm.email" /></td>
            <td><input v-model="editForm.role" /></td>
            <td>
              <button @click="saveEdit(user.ID)">Enregistrer</button>
              <button @click="cancelEdit">Annuler</button>
            </td>
          </template>
          <template v-else>
            <td>{{ user.nom }}</td>
            <td>{{ user.prenom }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.role }}</td>
            <td>
              <button @click="startEdit(user)">Modifier</button>
              <button @click="deleteUser(user.ID)">Supprimer</button>
            </td>
          </template>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
div {
  max-width: 900px;
  margin: 3rem auto;
  padding: 0 2rem;
}

h1 {
  font-family: var(--font-heading);
  font-weight: 600;
  color: var(--color-heading);
  font-size: 1.6rem;
  margin-bottom: 1.5rem;
}

table {
  width: 100%;
  table-layout: fixed;
  border-collapse: collapse;
}

th,
td {
  padding: 0.7rem;
  text-align: left;
  border-bottom: 1px solid var(--color-border);
  font-family: var(--font-body);
  font-size: 0.9rem;
}

th {
  font-family: var(--font-heading);
  font-weight: 600;
  color: var(--color-heading);
  font-size: 0.85rem;
}

th:last-child,
td:last-child {
  width: 190px;
}

td input {
  width: 100%;
  padding: 0.4rem;
  border: 1px solid var(--color-border-hover);
  border-radius: 4px;
  font-family: var(--font-body);
  font-size: 0.9rem;
}

button {
  padding: 0.4rem 0.8rem;
  margin-right: 0.4rem;
  border: 1px solid var(--color-primary);
  border-radius: 4px;
  background-color: #ffffff;
  color: var(--color-primary);
  font-family: var(--font-body);
  font-size: 0.85rem;
  cursor: pointer;
}

button:hover {
  background-color: var(--color-primary);
  color: #ffffff;
}
</style>
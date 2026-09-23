<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useRouter, useRoute } from 'vue-router'

const email = ref('')
const password = ref('')
const errorMessage = ref('')

const authStore = useAuthStore()
const router = useRouter()
const route = useRoute()

async function handleSubmit() {
  errorMessage.value = ''

  try {
    await authStore.login(email.value, password.value)
    router.push(route.query.redirect || '/')
  } catch (error) {
    errorMessage.value = error.message
  }
}
</script>

<template>
  <div>
    <h1>Connexion</h1>

    <form @submit.prevent="handleSubmit">
      <div>
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" required />
      </div>

      <div>
        <label for="password">Mot de passe</label>
        <input id="password" v-model="password" type="password" required />
      </div>

      <p v-if="errorMessage">{{ errorMessage }}</p>

      <button type="submit">Se connecter</button>
    </form>
  </div>
</template>

<style scoped>
div {
  max-width: 380px;
  margin: 4rem auto;
  padding: 2rem;
  border: 1px solid var(--color-border);
  border-radius: 8px;
}

h1 {
  font-family: var(--font-heading);
  font-weight: 600;
  color: var(--color-heading);
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
}

form div {
  border: none;
  padding: 0;
  margin: 0 0 1rem 0;
}

label {
  display: block;
  font-family: var(--font-body);
  font-size: 0.9rem;
  margin-bottom: 0.4rem;
}

input {
  width: 100%;
  padding: 0.6rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  font-family: var(--font-body);
  font-size: 1rem;
}

input:focus {
  outline: none;
  border-color: var(--color-border-hover);
}

button {
  width: 100%;
  padding: 0.7rem;
  background-color: var(--color-primary);
  color: #ffffff;
  border: none;
  border-radius: 4px;
  font-family: var(--font-body);
  font-weight: 500;
  font-size: 1rem;
  cursor: pointer;
}

button:hover {
  background-color: var(--color-accent);
}

p {
  color: #c0392b;
  font-size: 0.9rem;
  margin-bottom: 1rem;
}
</style>
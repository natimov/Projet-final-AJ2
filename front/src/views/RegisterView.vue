<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useI18n } from 'vue-i18n'

const nom = ref('')
const prenom = ref('')
const email = ref('')
const password = ref('')
const telephone = ref('')
const adresse = ref('')
const ville = ref('')
const errorMessage = ref('')

const authStore = useAuthStore()
const router = useRouter()
const { t } = useI18n()

async function handleSubmit() {
  errorMessage.value = ''

  try {
    await authStore.register({
      nom: nom.value,
      prenom: prenom.value,
      email: email.value,
      password: password.value,
      telephone: telephone.value,
      adresse: adresse.value,
      ville: ville.value,
    })
    router.push('/login')
  } catch {
    errorMessage.value = t('auth.registerError')
  }
}
</script>

<template>
  <div>
    <h1>{{ t('auth.registerTitle') }}</h1>

    <form @submit.prevent="handleSubmit">
      <div>
        <label for="nom">{{ t('auth.lastName') }}</label>
        <input id="nom" v-model="nom" type="text" required />
      </div>

      <div>
        <label for="prenom">{{ t('auth.firstName') }}</label>
        <input id="prenom" v-model="prenom" type="text" required />
      </div>

      <div>
        <label for="email">{{ t('auth.email') }}</label>
        <input id="email" v-model="email" type="email" required />
      </div>

      <div>
        <label for="password">{{ t('auth.password') }}</label>
        <input id="password" v-model="password" type="password" required />
      </div>

      <div>
        <label for="telephone">{{ t('auth.phone') }}</label>
        <input id="telephone" v-model="telephone" type="tel" />
      </div>

      <div>
        <label for="adresse">{{ t('auth.address') }}</label>
        <input id="adresse" v-model="adresse" type="text" />
      </div>

      <div>
        <label for="ville">{{ t('auth.city') }}</label>
        <input id="ville" v-model="ville" type="text" />
      </div>

      <p v-if="errorMessage">{{ errorMessage }}</p>

      <button type="submit">{{ t('auth.registerButton') }}</button>
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

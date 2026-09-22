<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

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
  } catch (error) {
    errorMessage.value = error.message
  }
}
</script>

<template>
  <div>
    <h1>Créer un compte</h1>

    <form @submit.prevent="handleSubmit">
      <div>
        <label for="nom">Nom</label>
        <input id="nom" v-model="nom" type="text" required />
      </div>

      <div>
        <label for="prenom">Prénom</label>
        <input id="prenom" v-model="prenom" type="text" required />
      </div>

      <div>
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" required />
      </div>

      <div>
        <label for="password">Mot de passe</label>
        <input id="password" v-model="password" type="password" required />
      </div>

      <div>
        <label for="telephone">Téléphone</label>
        <input id="telephone" v-model="telephone" type="tel" />
      </div>

      <div>
        <label for="adresse">Adresse</label>
        <input id="adresse" v-model="adresse" type="text" />
      </div>

      <div>
        <label for="ville">Ville</label>
        <input id="ville" v-model="ville" type="text" />
      </div>

      <p v-if="errorMessage">{{ errorMessage }}</p>

      <button type="submit">Créer mon compte</button>
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
<script setup>
import { RouterLink } from 'vue-router'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>

<template>
  <main class="home-page">
    <section class="hero">
      <div>
        <p class="eyebrow">{{ t('home.eyebrow') }}</p>
        <h1>{{ t('home.title') }}</h1>
        <p class="intro">{{ t('home.intro') }}</p>
        <div class="actions">
          <RouterLink class="button primary" to="/register">{{ t('home.createAccount') }}</RouterLink>
          <RouterLink class="button secondary" to="/login">{{ t('home.login') }}</RouterLink>
        </div>
      </div>

      <div class="illustration" aria-hidden="true">
        <div class="illustration-card">
          <span class="icon">↗</span>
          <p>{{ t('home.cardLabel') }}</p>
          <strong>{{ t('home.cardTitle') }}</strong>
          <div class="tags">
            <span>{{ t('home.repair') }}</span>
            <span>{{ t('home.reuse') }}</span>
            <span>{{ t('home.share') }}</span>
          </div>
        </div>
      </div>
    </section>

    <section class="discover">
      <p class="eyebrow">{{ t('home.projectLabel') }}</p>
      <h2>{{ t('home.projectTitle') }}</h2>
      <p class="section-intro">{{ t('home.projectIntro') }}</p>

      <div class="features">
        <article class="feature-card">
          <span>01</span>
          <h3>{{ t('home.giveTitle') }}</h3>
          <p>{{ t('home.giveText') }}</p>
        </article>
        <article class="feature-card">
          <span>02</span>
          <h3>{{ t('home.learnTitle') }}</h3>
          <p>{{ t('home.learnText') }}</p>
        </article>
        <article class="feature-card">
          <span>03</span>
          <h3>{{ t('home.impactTitle') }}</h3>
          <p>{{ t('home.impactText') }}</p>
        </article>
      </div>
    </section>

    <section class="community">
      <p class="eyebrow">{{ t('home.communityLabel') }}</p>
      <h2>{{ t('home.communityTitle') }}</h2>
      <p class="section-intro">{{ t('home.communityIntro') }}</p>

      <div class="features">
        <article class="feature-card actor-card">
          <span class="actor-icon" aria-hidden="true">01</span>
          <h3>{{ t('home.individualsTitle') }}</h3>
          <p>{{ t('home.individualsText') }}</p>
        </article>
        <article class="feature-card actor-card">
          <span class="actor-icon" aria-hidden="true">02</span>
          <h3>{{ t('home.craftspeopleTitle') }}</h3>
          <p>{{ t('home.craftspeopleText') }}</p>
        </article>
        <article class="feature-card actor-card">
          <span class="actor-icon" aria-hidden="true">03</span>
          <h3>{{ t('home.teamsTitle') }}</h3>
          <p>{{ t('home.teamsText') }}</p>
        </article>
      </div>
    </section>

    <section class="join">
      <div>
        <p class="eyebrow">{{ t('home.welcome') }}</p>
        <h2>{{ t('home.joinTitle') }}</h2>
        <p>{{ t('home.joinText') }}</p>
      </div>
      <RouterLink class="button primary" to="/register">{{ t('home.joinButton') }}</RouterLink>
    </section>
  </main>
</template>

<style scoped>
.home-page {
  width: 100%;
  padding: 2rem clamp(1rem, 5vw, 5rem) 4rem;
  color: #26372d;
  background: #f8f8f3;
}

.hero {
  max-width: 1480px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  align-items: center;
  gap: 5rem;
  min-height: 520px;
  margin: 0 auto;
  padding: 3rem 0 5rem;
}

.eyebrow {
  color: #607a52;
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.12em;
}

h1,
h2,
h3,
p {
  margin: 0;
}

h1,
h2,
h3,
strong {
  color: #244b38;
  font-family: var(--font-heading);
}

h1 {
  margin-top: 1rem;
  font-size: clamp(2.5rem, 5vw, 4rem);
  font-weight: 700;
  line-height: 1.12;
}

.intro,
.section-intro,
.join > div > p:last-child {
  color: #637168;
  line-height: 1.7;
}

.intro {
  max-width: 520px;
  margin-top: 1rem;
  font-size: 1.05rem;
}

.actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  margin-top: 1.5rem;
}

.button {
  display: inline-block;
  padding: 0.75rem 1.2rem;
  border: 1px solid #244b38;
  border-radius: 30px;
  font-weight: 600;
  text-align: center;
  text-decoration: none;
}

.primary {
  color: white;
  background: #244b38;
}

.secondary {
  color: #244b38;
  background: transparent;
}

.button:hover {
  background: #466d4f;
  color: white;
}

.button:focus-visible {
  outline: 3px solid #a2b86b;
  outline-offset: 3px;
}

.illustration {
  display: grid;
  min-height: 420px;
  place-items: center;
  border-radius: 45% 55% 55% 45%;
  background: #e8eddf;
}

.illustration-card {
  width: 270px;
  max-width: 80%;
  padding: 1.5rem;
  border-radius: 16px;
  background: #fffefa;
  box-shadow: 0 12px 30px #244b3820;
}

.icon {
  display: grid;
  width: 40px;
  height: 40px;
  place-items: center;
  border-radius: 50%;
  background: #e6edce;
  color: #244b38;
  font-size: 1.5rem;
}

.illustration-card p {
  margin-top: 1rem;
  color: #637168;
  font-size: 0.85rem;
}

.illustration-card strong {
  display: block;
  margin-top: 0.25rem;
  font-size: 1.5rem;
  line-height: 1.3;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.4rem;
  margin-top: 1rem;
}

.tags span {
  padding: 0.25rem 0.5rem;
  border-radius: 20px;
  background: #edf1e6;
  color: #52664e;
  font-size: 0.7rem;
}

.discover {
  max-width: 1480px;
  margin: 0 auto;
  padding: 4rem 0;
  border-top: 1px solid #e1e5db;
}

h2 {
  margin-top: 0.5rem;
  font-size: clamp(1.7rem, 3vw, 2.2rem);
  font-weight: 700;
}

.section-intro {
  max-width: 650px;
  margin-top: 0.6rem;
}

.features {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1rem;
  margin-top: 1.75rem;
}

.community {
  max-width: 1480px;
  margin: 0 auto;
  padding: 3rem 0 4rem;
  border-top: 1px solid #e1e5db;
}

.feature-card {
  min-height: 180px;
  padding: 1.3rem;
  border: 1px solid #e1e5db;
  border-radius: 14px;
  background: white;
}

.feature-card > span {
  color: #839267;
  font-size: 0.8rem;
  font-weight: 700;
}

.actor-icon {
  display: inline-grid;
  width: 38px;
  height: 38px;
  place-items: center;
  border-radius: 50%;
  background: #e8eddf;
}

.feature-card h3 {
  margin-top: 0.9rem;
  font-size: 1.1rem;
  font-weight: 700;
}

.feature-card p {
  margin-top: 0.4rem;
  color: #637168;
  line-height: 1.6;
}

.join {
  max-width: 1480px;
  min-height: 170px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1.5rem;
  padding: 1.75rem;
  border-radius: 16px;
  background: #e8eddf;
}

.join > div > p:last-child {
  margin-top: 0.5rem;
}

@media (max-width: 750px) {
  .hero {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .illustration {
    min-height: 280px;
  }

  .features {
    grid-template-columns: 1fr;
  }

  .feature-card {
    min-height: auto;
  }
}

@media (max-width: 550px) {
  .home-page {
    padding: 1.5rem 1rem 2.5rem;
  }

  .actions,
  .join {
    align-items: stretch;
    flex-direction: column;
  }

  .button {
    width: 100%;
  }

  .illustration {
    min-height: 250px;
  }
}
</style>

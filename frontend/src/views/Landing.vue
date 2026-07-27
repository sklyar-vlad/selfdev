<template>
  <div class="main-layout">
    <LandingHeader />
  </div>

  <main class="landing-page">
    <section class="hero">
      <div class="hero-content">
        <p class="eyebrow">SelfDev</p>
        <h1>Track everything that matters.</h1>
        <p class="hero-copy">
          A clean self-development tracker for habits, tasks, and goals. Create your account, verify
          your email, and keep your routine in one focused place.
        </p>

        <div class="hero-actions">
          <button @click="casdoorSignup" class="btn btn-primary">Create account</button>
        </div>
      </div>

      <div class="hero-panel">
        <div class="stat-card">
          <span class="stat-label">Flow</span>
          <strong>Register, verify, return</strong>
          <p>Simple onboarding built around email verification and secure sign-in.</p>
        </div>

        <div class="stat-card stat-card-accent">
          <span class="stat-label">Focus</span>
          <strong>Less clutter, more consistency</strong>
          <p>A calm interface with theme support, soft glass cards, and quick feedback.</p>
        </div>
      </div>

      <div class="hero-decoration"></div>
    </section>

    <section id="about" class="content-section">
      <div class="section-head">
        <p class="eyebrow">About</p>
        <h2>Built for daily progress without the noise.</h2>
      </div>

      <div class="section-card section-card-large">
        <p>
          SelfDev is a personal tracker designed to stay out of your way. The current product
          centers on account creation, email verification, secure login, and a private profile area
          so you can return to the same workspace every day.
        </p>
        <p>
          The tone is intentionally minimal: clear actions, calm visuals, and a lightweight flow
          that supports repeat use rather than constant setup.
        </p>
      </div>
    </section>

    <section id="features" class="content-section">
      <div class="section-head">
        <p class="eyebrow">Features</p>
        <h2>Everything the current build actually does.</h2>
      </div>

      <div class="feature-grid">
        <article v-for="feature in features" :key="feature.title" class="feature-card">
          <h3>{{ feature.title }}</h3>
          <p>{{ feature.text }}</p>
        </article>
      </div>
    </section>

    <section id="tech" class="content-section">
      <div class="section-head">
        <p class="eyebrow">Tech Stack</p>
        <h2>Modern tools, kept deliberately small.</h2>
      </div>

      <div class="tech-grid">
        <article v-for="group in techStack" :key="group.title" class="section-card tech-card">
          <h3>{{ group.title }}</h3>
          <p>{{ group.text }}</p>
        </article>
      </div>
    </section>

    <section id="faq" class="content-section">
      <div class="section-head">
        <p class="eyebrow">FAQ</p>
        <h2>Common questions, answered plainly.</h2>
      </div>

      <div class="faq-grid">
        <article v-for="item in faqs" :key="item.question" class="faq-card">
          <h3>{{ item.question }}</h3>
          <p>{{ item.answer }}</p>
        </article>
      </div>
    </section>

    <section class="cta-section">
      <div class="cta-card">
        <p class="eyebrow">Final CTA</p>
        <h2>Start your routine with less friction.</h2>
        <p>
          Create your account, verify your email, and keep your habits, tasks, and goals in one calm
          place.
        </p>

        <div class="hero-actions">
          <button @click="casdoorSignup" class="btn btn-primary">Create account</button>
        </div>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import LandingHeader from '@/components/Header/LandingHeader.vue'
import { config } from '@/config/env'

const casdoorSignup = () => {
  const url =
    `${config.authUrl}/signup/oauth/authorize` +
    `?client_id=${config.casdoorClientId}` +
    `&response_type=code` +
    `&scope=${encodeURIComponent('openid profile email')}` +
    `&redirect_uri=${encodeURIComponent(config.redirectUri)}`

  window.location.href = url
}

const features = [
  {
    title: 'Habit, task, and goal tracking',
    text: 'The product is positioned as a personal tracker for everyday self-development.',
  },
  {
    title: 'Username or email login',
    text: 'Users can sign in with either identifier, which keeps the login form simple.',
  },
]

const techStack = [
  {
    title: 'Frontend',
    text: 'Vue 3, Vue Router, TypeScript, Vite, and Vue Toastification.',
  },
  {
    title: 'Backend',
    text: 'Go 1.26.2 with JWT, pgx, Redis, bcrypt, Zap, and the Resend email API.',
  },
  {
    title: 'Data and delivery',
    text: 'PostgreSQL 18, Redis, Nginx, Docker, docker-compose, and Task.',
  },
]

const faqs = [
  {
    question: 'What is SelfDev?',
    answer: 'A personal tracker built around habits, tasks, and goals.',
  },
]
</script>

<style scoped>
/* =========================
   PAGE SHELL
========================= */

.landing-page {
  width: 100%;
  padding-bottom: 40px;
}

/* =========================
   SHARED LAYOUT
========================= */

.content-section,
.cta-section,
.hero {
  position: relative;
  width: min(1120px, calc(100% - 32px));
  margin: 0 auto;
}

.content-section {
  padding: 90px 0 0;
  scroll-margin-top: 96px;
}

.section-head {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 22px;
}

.eyebrow {
  font-size: 12px;
  font-weight: 700;

  letter-spacing: 0.14em;
  text-transform: uppercase;

  color: var(--accent-primary);
}

.section-head h2,
.cta-card h2 {
  font-size: clamp(28px, 3.4vw, 44px);

  line-height: 1.1;

  color: var(--text-primary);
}

/* =========================
   HERO
========================= */
.hero {
  min-height: auto;

  display: grid;

  grid-template-columns:
    minmax(0, 1.2fr)
    minmax(280px, 0.8fr);

  gap: 24px;

  align-items: center;

  padding: 120px 0 60px;
}

.hero-content,
.hero-panel {
  position: relative;

  z-index: 1;
}

.hero-content {
  display: flex;

  flex-direction: column;

  gap: 16px;

  padding: clamp(24px, 3vw, 40px);
}

.hero-panel {
  display: grid;

  gap: 16px;

  padding: clamp(18px, 2vw, 24px);
}

/* =========================
   HERO TITLE
========================= */

.hero h1 {
  font-family: 'Montserrat', sans-serif;

  font-size: clamp(36px, 6vw, 64px);

  font-weight: 800;

  line-height: 1.1;

  color: var(--text-primary);

  text-shadow:
    0 0 10px rgba(149, 162, 223, 0.18),
    0 0 25px rgba(59, 130, 246, 0.12);
}

.hero-copy {
  max-width: 560px;

  font-size: clamp(15px, 1.4vw, 18px);

  color: var(--text-secondary);

  line-height: 1.6;
}

/* =========================
   GLASS FIX
========================= */

.hero-content,
.hero-panel,
.section-card,
.feature-card,
.faq-card,
.cta-card {
  background: var(--surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--surface-border);
  -webkit-backdrop-filter: blur(20px) !important;
  backdrop-filter: blur(20px) !important;
  box-shadow: var(--shadow-md);
}
/* =========================
   STAT CARDS
========================= */

.stat-card {
  padding: 18px;

  border-radius: 14px;

  border: 1px solid var(--border-default);

  background: var(--surface);
}

.stat-card-accent {
  background: linear-gradient(135deg, rgba(100, 200, 255, 0.14), rgba(108, 123, 255, 0.12));
}

.stat-label {
  display: inline-block;

  margin-bottom: 10px;

  font-size: 12px;

  font-weight: 700;

  letter-spacing: 0.12em;

  text-transform: uppercase;

  color: var(--accent-primary);
}

.stat-card strong,
.feature-card h3,
.tech-card h3,
.faq-card h3 {
  display: block;

  margin-bottom: 8px;

  font-size: 18px;

  color: var(--text-primary);
}

.stat-card p,
.section-card p,
.feature-card p,
.tech-card p,
.faq-card p,
.cta-card p {
  color: var(--text-secondary);

  line-height: 1.6;
}

/* =========================
   SECTION CARDS
========================= */

.section-card-large {
  padding: clamp(20px, 3vw, 30px);

  display: grid;

  gap: 12px;
}

/* =========================
   BUTTONS
========================= */

.hero-actions {
  display: flex;

  gap: 14px;

  flex-wrap: wrap;

  margin-top: 10px;
}

.btn {
  padding: 12px 24px;
  border: none;
  border-radius: 10px;
  font-weight: 700;

  display: inline-flex;

  align-items: center;

  justify-content: center;
  transition: all 0.25s ease;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-dark));

  color: white;

  box-shadow: 0 10px 30px rgba(59, 130, 246, 0.25);
}

.btn-primary:hover {
  transform: translateY(-3px);
}

/* =========================
   GRIDS
========================= */

.feature-grid,
.faq-grid,
.tech-grid {
  display: grid;

  gap: 16px;
}

.feature-grid,
.tech-grid {
  grid-template-columns: repeat(3, minmax(0, 1fr));
}

.faq-grid {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.feature-card,
.faq-card,
.tech-card {
  padding: 20px;
}

/* =========================
   CTA
========================= */

.cta-section {
  padding: 100px 0 40px;
}

.cta-card {
  padding: clamp(24px, 3vw, 36px);

  display: flex;

  flex-direction: column;

  gap: 14px;

  align-items: flex-start;
}

/* =========================
   HERO GLOW
========================= */

.hero::before {
  content: '';

  position: absolute;

  width: 600px;

  height: 600px;

  background: radial-gradient(circle, rgba(149, 162, 223, 0.12), transparent 60%);

  filter: blur(50px);

  top: 50%;

  left: 50%;

  transform: translate(-50%, -50%);

  z-index: 0;

  pointer-events: none;
}

/* =========================
   RESPONSIVE
========================= */

@media (max-width: 1024px) {
  .hero {
    grid-template-columns: 1fr;

    padding-top: 80px;
  }

  .feature-grid,
  .tech-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 768px) {
  .content-section,
  .cta-section,
  .hero {
    width: calc(100% - 24px);
  }

  .hero {
    min-height: auto;

    padding: 96px 0 30px;
  }

  .hero-actions {
    flex-direction: column;

    width: 100%;
  }

  .btn {
    width: 100%;
  }

  .feature-grid,
  .tech-grid,
  .faq-grid {
    grid-template-columns: 1fr;
  }
}
</style>

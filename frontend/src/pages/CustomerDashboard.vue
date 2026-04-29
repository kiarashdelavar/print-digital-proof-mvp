<script setup>
import { onMounted, ref, computed } from "vue";
import { RouterLink } from "vue-router";
import { getProofs } from "../services/proofService";
import printLogo from "../assets/print-logo.svg";

const proofs = ref([]);
const isLoading = ref(true);
const errorMessage = ref("");

const pendingCount = computed(() =>
  proofs.value.filter((proof) => proof.status === "WAITING_FOR_APPROVAL").length
);

const approvedCount = computed(() =>
  proofs.value.filter((proof) => proof.status === "APPROVED").length
);

const changesCount = computed(() =>
  proofs.value.filter((proof) => proof.status === "CHANGES_REQUESTED").length
);

function getStatusLabel(status) {
  if (status === "WAITING_FOR_APPROVAL") return "Waiting for approval";
  if (status === "APPROVED") return "Approved";
  if (status === "CHANGES_REQUESTED") return "Changes requested";
  return status;
}

function getStatusClass(status) {
  if (status === "APPROVED") return "status-approved";
  if (status === "CHANGES_REQUESTED") return "status-changes";
  return "status-waiting";
}

onMounted(async () => {
  try {
    proofs.value = await getProofs();
  } catch (error) {
    errorMessage.value = "Could not load proofs. Please check if the Go backend is running.";
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <main class="page">
    <header class="top-nav">
      <div class="brand">
        <img :src="printLogo" alt="Print.com logo" class="brand-logo" />

       <nav class="nav-links">
  <RouterLink to="/customer/dashboard">Customer dashboard</RouterLink>
  <RouterLink to="/file-checker/dashboard">File checker</RouterLink>
  <RouterLink to="/file-checker/upload">Upload proof</RouterLink>
  <a href="#">Catalogue</a>
</nav>
      </div>

      <div class="nav-actions">
        <button class="language-button">🌐 English</button>
        <button class="login-button">Login</button>
        <button class="register-button">Register</button>
      </div>
    </header>

    <section class="hero-card">
      <div class="hero-content">
        <p class="eyebrow">
          Designed for <span>proof reviewers*</span>
        </p>

        <h1>
          Your first-class<br />
          proof approval<br />
          workflow
        </h1>

        <p class="hero-text">
          Review digital proofs, approve production files, and request changes
          in one clear customer workflow.
        </p>

       <div class="hero-actions">
  <RouterLink to="/customer/proofs/1" class="hero-button">
    Review first proof
    <span>↗</span>
  </RouterLink>

  <RouterLink to="/file-checker/dashboard" class="hero-button secondary-hero-button">
    File checker view
    <span>↗</span>
  </RouterLink>
</div>
      </div>

      <div class="hero-visual">
        <div class="window-card">
          <div class="window-top">
            <span></span>
            <span></span>
            <span></span>
          </div>

          <div class="window-content">
            <div class="proof-mini-header">
              <div>
                <p>Digital proof</p>
                <strong>Business Cards</strong>
              </div>
              <span>PDF</span>
            </div>

            <div class="proof-preview-shape">
              <div class="paper paper-one"></div>
              <div class="paper paper-two"></div>
            </div>

            <div class="mini-actions">
              <div class="mini-button pink"></div>
              <div class="mini-button dark"></div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="intro-panel">
      <div class="intro-title">
  <h2>Print management for professionals</h2>
  <p>Your proof approval gateway for faster production decisions.</p>

  <RouterLink to="/file-checker/upload" class="upload-shortcut">
    Upload a new proof
    <span>↗</span>
  </RouterLink>
</div>

      <div class="feature-row">
        <div class="feature-item">
          <span class="feature-icon pink-icon">✦</span>
          <strong>Get proof-ready</strong>
        </div>

        <div class="feature-item">
          <span class="feature-icon blue-icon">⏱</span>
          <strong>Approve fast</strong>
        </div>

        <div class="feature-item">
          <span class="feature-icon yellow-icon">▣</span>
          <strong>Collaborate smartly</strong>
        </div>

        <div class="feature-item">
          <span class="feature-icon green-icon">✓</span>
          <strong>Track clearly</strong>
        </div>
      </div>
    </section>

    <section class="stats-grid">
      <article class="stat-card">
        <p>Waiting</p>
        <strong>{{ pendingCount }}</strong>
      </article>

      <article class="stat-card">
        <p>Approved</p>
        <strong>{{ approvedCount }}</strong>
      </article>

      <article class="stat-card">
        <p>Changes requested</p>
        <strong>{{ changesCount }}</strong>
      </article>
    </section>

    <section class="content-card">
      <div class="section-header">
        <div>
          <p class="section-eyebrow">Customer workspace</p>
          <h2>Proofs waiting for review</h2>
          <p>These proofs are loaded from the Go backend API.</p>
        </div>
      </div>

      <p v-if="isLoading" class="muted">Loading proofs...</p>

      <p v-else-if="errorMessage" class="error">
        {{ errorMessage }}
      </p>

      <div v-else class="proof-grid">
        <article v-for="proof in proofs" :key="proof.id" class="proof-card">
          <div class="proof-top">
            <div>
              <p class="order-id">{{ proof.orderId }}</p>
              <h3>{{ proof.productName }}</h3>
            </div>

            <span :class="['status-pill', getStatusClass(proof.status)]">
              {{ getStatusLabel(proof.status) }}
            </span>
          </div>

          <div class="proof-body">
            <p><strong>Customer:</strong> {{ proof.customerName }}</p>
            <p><strong>File:</strong> {{ proof.fileName }}</p>
            <p v-if="proof.comment"><strong>Comment:</strong> {{ proof.comment }}</p>
          </div>

          <div class="proof-footer">
            <RouterLink :to="`/customer/proofs/${proof.id}`" class="primary-button">
              Review proof
              <span>↗</span>
            </RouterLink>
          </div>
        </article>
      </div>
    </section>
  </main>
</template>

<style scoped>
.page {
  min-height: 100vh;
  background: #f4f1f1;
  padding: 0 28px 48px;
  color: #071733;
  font-family: Inter, Arial, sans-serif;
}

.top-nav {
  height: 92px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1680px;
  margin: 0 auto;
}

.brand {
  display: flex;
  align-items: center;
  gap: 28px;
}

.brand-logo {
  width: 76px;
  height: 76px;
  object-fit: contain;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 28px;
}

.nav-links a {
  color: #26364d;
  text-decoration: none;
  font-size: 15px;
  font-weight: 700;
}

.nav-links a::after {
  content: "⌄";
  margin-left: 7px;
  font-size: 12px;
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.language-button,
.login-button,
.register-button {
  border: 0;
  border-radius: 10px;
  padding: 12px 18px;
  font-weight: 900;
  cursor: pointer;
}

.language-button {
  background: transparent;
  color: #26364d;
}

.login-button {
  background: white;
  color: #317c95;
  border: 2px solid #a8d6e1;
}

.register-button {
  background: #4f9bb1;
  color: white;
}

.hero-card {
  max-width: 1680px;
  min-height: 580px;
  margin: 0 auto 28px;
  border-radius: 24px;
  background:
    radial-gradient(circle at 72% 25%, rgba(255, 255, 255, 0.55), transparent 22%),
    linear-gradient(135deg, #78d0e4 0%, #7fd3e8 48%, #bfeef6 100%);
  display: grid;
  grid-template-columns: 1fr 1.15fr;
  overflow: hidden;
  position: relative;
}

.hero-content {
  padding: 110px 0 90px 280px;
  position: relative;
  z-index: 2;
}

.eyebrow {
  margin: 0 0 24px;
  font-size: 18px;
  font-weight: 900;
}

.eyebrow span {
  font-style: italic;
  text-decoration: underline;
}

h1 {
  margin: 0;
  color: white;
  font-size: 64px;
  line-height: 1.14;
  letter-spacing: -1px;
  text-transform: uppercase;
  max-width: 520px;
  font-weight: 950;
}

.hero-text {
  max-width: 480px;
  margin: 24px 0 0;
  color: #0c2538;
  font-size: 17px;
  line-height: 1.5;
  font-weight: 650;
}

.hero-button {
  margin-top: 34px;
  display: inline-flex;
  align-items: center;
  gap: 14px;
  background: #ff0050;
  color: white;
  text-decoration: none;
  border-radius: 10px;
  padding: 18px 28px;
  font-weight: 950;
  box-shadow: 0 16px 30px rgba(255, 0, 80, 0.25);
}

.hero-button span {
  font-size: 22px;
}

.hero-visual {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-right: 120px;
}

.window-card {
  width: 520px;
  border-radius: 30px;
  background: white;
  box-shadow: 0 30px 70px rgba(7, 23, 51, 0.25);
  overflow: hidden;
  transform: rotate(-2deg);
}

.window-top {
  height: 54px;
  background: #f4f1f1;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0 20px;
}

.window-top span {
  width: 12px;
  height: 12px;
  border-radius: 999px;
  background: #ff0050;
}

.window-top span:nth-child(2) {
  background: #ffd166;
}

.window-top span:nth-child(3) {
  background: #7fd3e8;
}

.window-content {
  padding: 30px;
}

.proof-mini-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.proof-mini-header p {
  margin: 0 0 6px;
  color: #607085;
  font-weight: 800;
}

.proof-mini-header strong {
  font-size: 24px;
}

.proof-mini-header span {
  background: #fff3c4;
  color: #8a5b00;
  border-radius: 999px;
  padding: 8px 12px;
  font-weight: 950;
}

.proof-preview-shape {
  height: 220px;
  background: #eaf9fc;
  border-radius: 22px;
  margin: 28px 0;
  position: relative;
  overflow: hidden;
}

.paper {
  position: absolute;
  border-radius: 18px;
  background: white;
  box-shadow: 0 14px 30px rgba(7, 23, 51, 0.12);
}

.paper-one {
  width: 150px;
  height: 110px;
  left: 82px;
  top: 52px;
  transform: rotate(-12deg);
}

.paper-two {
  width: 180px;
  height: 125px;
  right: 70px;
  top: 46px;
  transform: rotate(8deg);
}

.mini-actions {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
}

.mini-button {
  height: 46px;
  border-radius: 12px;
}

.mini-button.pink {
  background: #ff0050;
}

.mini-button.dark {
  background: #111827;
}

.intro-panel {
  max-width: 1340px;
  margin: 0 auto 36px;
  background: white;
  border-radius: 0 0 26px 26px;
  padding: 44px 80px 54px;
  text-align: center;
}

.intro-title h2 {
  margin: 0;
  color: #ff0050;
  font-size: 32px;
}

.intro-title p {
  margin: 16px 0 38px;
  color: #667085;
  font-size: 17px;
}

.feature-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 34px;
}

.feature-item {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 14px;
  font-size: 18px;
}

.feature-icon {
  display: inline-grid;
  place-items: center;
  width: 48px;
  height: 48px;
  border-radius: 14px;
  font-weight: 950;
}

.pink-icon {
  background: #ffe0e8;
}

.blue-icon {
  background: #d9f7ff;
}

.yellow-icon {
  background: #fff3c4;
}

.green-icon {
  background: #d9fbe7;
}

.stats-grid {
  max-width: 1680px;
  margin: 0 auto 28px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 22px;
}

.stat-card {
  background: white;
  border-radius: 24px;
  padding: 28px;
  box-shadow: 0 12px 26px rgba(7, 23, 51, 0.04);
}

.stat-card p {
  margin: 0 0 8px;
  color: #667085;
  font-weight: 850;
}

.stat-card strong {
  font-size: 42px;
  color: #071733;
}

.content-card {
  max-width: 1680px;
  margin: 0 auto;
  background: white;
  border-radius: 28px;
  padding: 36px;
  box-shadow: 0 14px 34px rgba(7, 23, 51, 0.04);
}

.section-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 26px;
}

.section-eyebrow {
  margin: 0 0 8px;
  color: #ff0050;
  font-weight: 950;
}

.section-header h2 {
  margin: 0;
  font-size: 32px;
}

.section-header p {
  margin: 8px 0 0;
  color: #667085;
}

.proof-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(420px, 1fr));
  gap: 20px;
}

.proof-card {
  border: 1px solid #eee4e7;
  border-radius: 24px;
  padding: 24px;
  background: #fffafa;
  min-height: 210px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.proof-top {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  align-items: flex-start;
}

.order-id {
  color: #667085;
  font-weight: 850;
  margin: 0 0 6px;
}

h3 {
  margin: 0;
  font-size: 24px;
}

.proof-body {
  margin: 18px 0;
  color: #26364d;
}

.proof-body p {
  margin: 8px 0;
}

.status-pill {
  white-space: nowrap;
  border-radius: 999px;
  padding: 9px 14px;
  font-size: 13px;
  font-weight: 950;
}

.status-waiting {
  background: #fff3c4;
  color: #8a5b00;
}

.status-approved {
  background: #d9fbe7;
  color: #047857;
}

.status-changes {
  background: #ffe0e8;
  color: #be123c;
}

.primary-button {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  text-decoration: none;
  background: #ff0050;
  color: white;
  border-radius: 12px;
  padding: 14px 20px;
  font-weight: 950;
  box-shadow: 0 12px 24px rgba(255, 0, 80, 0.18);
}

.muted {
  color: #667085;
}

.error {
  color: #be123c;
  font-weight: 850;
}

@media (max-width: 1100px) {
  .nav-links {
    display: none;
  }

  .hero-card {
    grid-template-columns: 1fr;
  }

  .hero-content {
    padding: 70px 44px;
  }

  .hero-visual {
    padding: 0 44px 60px;
  }

  h1 {
    font-size: 48px;
  }

  .feature-row,
  .stats-grid {
    grid-template-columns: 1fr 1fr;
  }
}

@media (max-width: 720px) {
  .page {
    padding: 0 14px 32px;
  }

  .top-nav {
    height: auto;
    padding: 16px 0;
  }

  .nav-actions {
    display: none;
  }

  .hero-content {
    padding: 48px 28px;
  }

  .hero-visual {
    display: none;
  }

  h1 {
    font-size: 38px;
  }

  .intro-panel {
    padding: 32px 20px;
  }

  .feature-row,
  .stats-grid {
    grid-template-columns: 1fr;
  }

  .proof-grid {
    grid-template-columns: 1fr;
  }
}
</style>
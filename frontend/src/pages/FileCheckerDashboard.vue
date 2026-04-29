<script setup>
import { computed, onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import { getProofs } from "../services/proofService";
import printLogo from "../assets/print-logo.svg";

const proofs = ref([]);
const isLoading = ref(true);
const errorMessage = ref("");

const waitingCount = computed(() =>
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
          <RouterLink to="/file-checker/dashboard">File checker</RouterLink>
          <RouterLink to="/file-checker/upload">Upload proof</RouterLink>
          <RouterLink to="/customer/dashboard">Customer view</RouterLink>
        </nav>
      </div>

      <div class="nav-actions">
        <button class="language-button">🌐 English</button>
        <button class="login-button">Login</button>
        <button class="register-button">Register</button>
      </div>
    </header>

    <section class="hero-card">
      <div>
        <p class="eyebrow">Internal workspace</p>
        <h1>
          Manage digital<br />
          proof approvals
        </h1>
        <p class="hero-text">
          Upload customer proofs, track approval statuses, and keep production moving.
        </p>

        <RouterLink to="/file-checker/upload" class="hero-button">
          Upload new proof
          <span>↗</span>
        </RouterLink>
      </div>

      <div class="hero-panel">
        <p>Approval status</p>
        <strong>{{ proofs.length }} proofs</strong>
        <span>Loaded from Go API</span>
      </div>
    </section>

    <section class="stats-grid">
      <article class="stat-card">
        <p>Waiting</p>
        <strong>{{ waitingCount }}</strong>
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
          <p class="section-eyebrow">Production overview</p>
          <h2>All digital proofs</h2>
          <p>Track which proofs are waiting, approved, or need changes.</p>
        </div>

        <RouterLink to="/file-checker/upload" class="secondary-link">
          + Upload proof
        </RouterLink>
      </div>

      <p v-if="isLoading" class="muted">Loading proofs...</p>

      <p v-else-if="errorMessage" class="error">
        {{ errorMessage }}
      </p>

      <div v-else class="table-card">
        <div class="table-header">
          <span>Order</span>
          <span>Product</span>
          <span>Customer</span>
          <span>Status</span>
          <span>Customer link</span>
        </div>

        <div v-for="proof in proofs" :key="proof.id" class="table-row">
          <span>{{ proof.orderId }}</span>
          <strong>{{ proof.productName }}</strong>
          <span>{{ proof.customerName }}</span>

          <span :class="['status-pill', getStatusClass(proof.status)]">
            {{ getStatusLabel(proof.status) }}
          </span>

          <RouterLink :to="`/customer/proofs/${proof.id}`" class="review-link">
            Open review page
          </RouterLink>
        </div>
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
  font-weight: 800;
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
  min-height: 440px;
  margin: 0 auto 28px;
  border-radius: 28px;
  background:
    radial-gradient(circle at 75% 24%, rgba(255, 255, 255, 0.55), transparent 24%),
    linear-gradient(135deg, #78d0e4 0%, #bdeef6 100%);
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 70px 90px;
}

.eyebrow {
  margin: 0 0 24px;
  font-weight: 950;
}

h1 {
  margin: 0;
  color: white;
  font-size: 64px;
  line-height: 1.08;
  text-transform: uppercase;
  letter-spacing: -1px;
  font-weight: 950;
}

.hero-text {
  max-width: 520px;
  margin: 26px 0 0;
  color: #0c2538;
  font-size: 18px;
  line-height: 1.55;
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
  border-radius: 12px;
  padding: 18px 28px;
  font-weight: 950;
  box-shadow: 0 16px 30px rgba(255, 0, 80, 0.25);
}

.hero-panel {
  width: 360px;
  border-radius: 28px;
  background: white;
  padding: 34px;
  box-shadow: 0 30px 70px rgba(7, 23, 51, 0.18);
}

.hero-panel p {
  margin: 0 0 10px;
  color: #667085;
  font-weight: 850;
}

.hero-panel strong {
  display: block;
  font-size: 42px;
  margin-bottom: 8px;
}

.hero-panel span {
  color: #ff0050;
  font-weight: 900;
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
  gap: 20px;
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

.secondary-link {
  height: fit-content;
  background: #ff0050;
  color: white;
  text-decoration: none;
  border-radius: 12px;
  padding: 14px 20px;
  font-weight: 950;
}

.table-card {
  border: 1px solid #eee4e7;
  border-radius: 22px;
  overflow: hidden;
}

.table-header,
.table-row {
  display: grid;
  grid-template-columns: 1fr 1.4fr 1.4fr 1.2fr 1fr;
  gap: 18px;
  align-items: center;
  padding: 18px 22px;
}

.table-header {
  background: #fffafa;
  color: #667085;
  font-weight: 950;
}

.table-row {
  border-top: 1px solid #eee4e7;
}

.status-pill {
  width: fit-content;
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

.review-link {
  color: #ff0050;
  font-weight: 950;
  text-decoration: none;
}

.muted {
  color: #667085;
}

.error {
  color: #be123c;
  font-weight: 850;
}

@media (max-width: 950px) {
  .hero-card {
    flex-direction: column;
    align-items: flex-start;
    padding: 44px;
  }

  h1 {
    font-size: 42px;
  }

  .hero-panel {
    width: 100%;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  .table-header {
    display: none;
  }

  .table-row {
    grid-template-columns: 1fr;
  }

  .nav-actions {
    display: none;
  }
}
</style>
<script setup>
import { computed, onMounted, ref } from "vue";
import { useRoute, RouterLink } from "vue-router";
import { approveProof, getProofById, requestChanges } from "../services/proofService";
import printLogo from "../assets/print-logo.svg";

const route = useRoute();

const proof = ref(null);
const isLoading = ref(true);
const errorMessage = ref("");
const comment = ref("");
const successMessage = ref("");

const statusLabel = computed(() => {
  if (!proof.value) return "";

  if (proof.value.status === "WAITING_FOR_APPROVAL") return "Waiting for approval";
  if (proof.value.status === "APPROVED") return "Approved";
  if (proof.value.status === "CHANGES_REQUESTED") return "Changes requested";

  return proof.value.status;
});

const statusClass = computed(() => {
  if (!proof.value) return "status-waiting";

  if (proof.value.status === "APPROVED") return "status-approved";
  if (proof.value.status === "CHANGES_REQUESTED") return "status-changes";

  return "status-waiting";
});

async function loadProof() {
  try {
    proof.value = await getProofById(route.params.id);
    comment.value = proof.value.comment || "";
  } catch (error) {
    errorMessage.value = "Could not load proof.";
  } finally {
    isLoading.value = false;
  }
}

async function handleApprove() {
  successMessage.value = "";
  errorMessage.value = "";

  try {
    proof.value = await approveProof(route.params.id);
    comment.value = "";
    successMessage.value = "Proof approved successfully.";
  } catch (error) {
    errorMessage.value = "Could not approve proof.";
  }
}

async function handleRequestChanges() {
  successMessage.value = "";
  errorMessage.value = "";

  if (!comment.value.trim()) {
    errorMessage.value = "Please write a comment before requesting changes.";
    return;
  }

  try {
    proof.value = await requestChanges(route.params.id, comment.value);
    successMessage.value = "Changes requested successfully.";
  } catch (error) {
    errorMessage.value = "Could not request changes.";
  }
}

onMounted(loadProof);
</script>

<template>
  <main class="page">
    <header class="top-nav">
      <div class="brand">
        <img :src="printLogo" alt="Print.com logo" class="brand-logo" />

        <nav class="nav-links">
          <RouterLink to="/customer/dashboard">Dashboard</RouterLink>
          <a href="#">Proofs</a>
          <a href="#">Orders</a>
          <a href="#">Support</a>
        </nav>
      </div>

      <div class="nav-actions">
        <button class="language-button">🌐 English</button>
        <button class="login-button">Login</button>
        <button class="register-button">Register</button>
      </div>
    </header>

    <RouterLink to="/customer/dashboard" class="back-link">
      ← Back to dashboard
    </RouterLink>

    <p v-if="isLoading" class="loading-text">Loading proof...</p>

    <section v-else-if="proof" class="review-shell">
      <aside class="side-panel">
        <p class="eyebrow">Digital proof</p>

        <h1>
          Check your<br />
          production file
        </h1>

        <p class="side-text">
          Review the digital proof before production starts. Approve the file if
          everything is correct, or request changes with a clear note.
        </p>

        <div class="side-feature">
          <span>1</span>
          <p>Open proof file</p>
        </div>

        <div class="side-feature">
          <span>2</span>
          <p>Check content and layout</p>
        </div>

        <div class="side-feature">
          <span>3</span>
          <p>Approve or request changes</p>
        </div>
      </aside>

      <section class="review-card">
        <div class="review-header">
          <div>
            <p class="section-eyebrow">Proof review</p>
            <h2>{{ proof.productName }}</h2>
            <p class="muted">{{ proof.orderId }} · {{ proof.customerName }}</p>
          </div>

          <span :class="['status-pill', statusClass]">
            {{ statusLabel }}
          </span>
        </div>

        <div class="meta-grid">
          <article>
            <p>Customer</p>
            <strong>{{ proof.customerName }}</strong>
          </article>

          <article>
            <p>Order</p>
            <strong>{{ proof.orderId }}</strong>
          </article>

          <article>
            <p>File</p>
            <strong>{{ proof.fileName }}</strong>
          </article>
        </div>

        <div class="preview-box">
          <div class="preview-window">
            <div class="preview-top">
              <span></span>
              <span></span>
              <span></span>
            </div>

            <div class="preview-content">
              <p class="preview-title">Proof file preview</p>
              <p class="muted">{{ proof.fileName }}</p>

              <div class="paper-stack">
                <div class="paper paper-one"></div>
                <div class="paper paper-two"></div>
              </div>

              <p class="preview-note">
                File preview will be implemented in the upload/proof viewer task.
              </p>
            </div>
          </div>
        </div>

        <div class="action-area">
          <button class="approve-button" @click="handleApprove">
            Approve proof
            <span>✓</span>
          </button>

          <div class="changes-box">
            <label for="comment">Request changes</label>
            <textarea
              id="comment"
              v-model="comment"
              placeholder="Example: Please make the logo bigger."
            />

            <button class="secondary-button" @click="handleRequestChanges">
              Submit changes
              <span>↗</span>
            </button>
          </div>
        </div>

        <p v-if="successMessage" class="success">{{ successMessage }}</p>
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </section>
    </section>

    <p v-else class="error">{{ errorMessage }}</p>
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

.nav-links a::after {
  content: "⌄";
  margin-left: 7px;
  font-size: 12px;
}

.nav-links a:first-child::after {
  content: "";
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

.back-link {
  display: flex;
  max-width: 1680px;
  margin: 0 auto 20px;
  color: #071733;
  font-weight: 950;
  text-decoration: none;
}

.loading-text {
  max-width: 1680px;
  margin: 30px auto;
  font-weight: 800;
}

.review-shell {
  max-width: 1680px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 0.9fr 1.35fr;
  gap: 28px;
  align-items: stretch;
}

.side-panel {
  border-radius: 28px;
  background:
    radial-gradient(circle at 75% 25%, rgba(255, 255, 255, 0.55), transparent 24%),
    linear-gradient(135deg, #78d0e4 0%, #bdeef6 100%);
  padding: 58px;
  min-height: 720px;
}

.eyebrow {
  margin: 0 0 26px;
  font-weight: 950;
}

h1 {
  margin: 0;
  color: white;
  font-size: 68px;
  line-height: 1.08;
  text-transform: uppercase;
  letter-spacing: -1px;
  font-weight: 950;
}

.side-text {
  max-width: 520px;
  margin: 28px 0 34px;
  color: #0c2538;
  font-size: 18px;
  line-height: 1.55;
  font-weight: 650;
}

.side-feature {
  display: flex;
  align-items: center;
  gap: 14px;
  background: rgba(255, 255, 255, 0.65);
  border-radius: 18px;
  padding: 16px;
  margin-bottom: 14px;
  max-width: 430px;
}

.side-feature span {
  display: grid;
  place-items: center;
  width: 36px;
  height: 36px;
  border-radius: 999px;
  background: #ff0050;
  color: white;
  font-weight: 950;
}

.side-feature p {
  margin: 0;
  font-weight: 850;
}

.review-card {
  background: white;
  border-radius: 28px;
  padding: 38px;
  box-shadow: 0 14px 34px rgba(7, 23, 51, 0.04);
}

.review-header {
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

h2 {
  margin: 0;
  font-size: 46px;
  letter-spacing: -0.8px;
}

.muted {
  color: #667085;
}

.status-pill {
  white-space: nowrap;
  border-radius: 999px;
  padding: 12px 18px;
  font-size: 14px;
  font-weight: 950;
  height: fit-content;
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

.meta-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 14px;
  margin-bottom: 24px;
}

.meta-grid article {
  background: #fffafa;
  border: 1px solid #eee4e7;
  border-radius: 18px;
  padding: 18px;
}

.meta-grid p {
  margin: 0 0 6px;
  color: #667085;
  font-weight: 850;
}

.meta-grid strong {
  color: #071733;
}

.preview-box {
  border-radius: 26px;
  background: #eaf9fc;
  border: 1px dashed #7fd3e8;
  padding: 24px;
  margin-bottom: 26px;
}

.preview-window {
  max-width: 680px;
  margin: 0 auto;
  border-radius: 24px;
  background: white;
  overflow: hidden;
  box-shadow: 0 18px 40px rgba(7, 23, 51, 0.08);
}

.preview-top {
  height: 44px;
  background: #f4f1f1;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0 18px;
}

.preview-top span {
  width: 11px;
  height: 11px;
  border-radius: 999px;
  background: #ff0050;
}

.preview-top span:nth-child(2) {
  background: #ffd166;
}

.preview-top span:nth-child(3) {
  background: #7fd3e8;
}

.preview-content {
  padding: 34px;
  text-align: center;
}

.preview-title {
  margin: 0;
  font-size: 28px;
  font-weight: 950;
}

.paper-stack {
  height: 160px;
  position: relative;
  margin: 20px auto;
  max-width: 360px;
  background: #eaf9fc;
  border-radius: 20px;
  overflow: hidden;
}

.paper {
  position: absolute;
  background: white;
  border-radius: 16px;
  box-shadow: 0 12px 26px rgba(7, 23, 51, 0.1);
}

.paper-one {
  width: 135px;
  height: 95px;
  left: 72px;
  top: 34px;
  transform: rotate(-13deg);
}

.paper-two {
  width: 145px;
  height: 100px;
  right: 68px;
  top: 34px;
  transform: rotate(8deg);
}

.preview-note {
  color: #26364d;
  font-weight: 650;
}

.action-area {
  display: grid;
  grid-template-columns: 260px 1fr;
  gap: 24px;
  align-items: start;
}

.approve-button,
.secondary-button {
  border: none;
  cursor: pointer;
  border-radius: 14px;
  padding: 16px 20px;
  font-weight: 950;
}

.approve-button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  background: #ff0050;
  color: white;
  box-shadow: 0 12px 24px rgba(255, 0, 80, 0.18);
}

.secondary-button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  background: #111827;
  color: white;
  margin-top: 12px;
}

.changes-box {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

label {
  font-weight: 950;
}

textarea {
  min-height: 126px;
  border-radius: 16px;
  border: 1px solid #d1d5db;
  padding: 16px;
  font-family: inherit;
  resize: vertical;
}

textarea:focus {
  outline: 3px solid rgba(127, 211, 232, 0.35);
  border-color: #7fd3e8;
}

.success {
  color: #047857;
  font-weight: 900;
  margin-top: 22px;
}

.error {
  color: #be123c;
  font-weight: 900;
  margin-top: 22px;
}

@media (max-width: 1200px) {
  .review-shell {
    grid-template-columns: 1fr;
  }

  .side-panel {
    min-height: auto;
  }

  h1 {
    font-size: 48px;
  }
}

@media (max-width: 760px) {
  .page {
    padding: 0 14px 32px;
  }

  .top-nav {
    height: auto;
    padding: 16px 0;
  }

  .nav-links,
  .nav-actions {
    display: none;
  }

  .side-panel {
    padding: 32px;
  }

  h1 {
    font-size: 38px;
  }

  h2 {
    font-size: 34px;
  }

  .review-card {
    padding: 24px;
  }

  .review-header,
  .action-area {
    grid-template-columns: 1fr;
    display: grid;
  }

  .meta-grid {
    grid-template-columns: 1fr;
  }
}
</style>
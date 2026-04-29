<script setup>
import { onMounted, ref } from "vue";
import { useRoute, RouterLink } from "vue-router";
import { approveProof, getProofById, requestChanges } from "../services/proofService";

const route = useRoute();

const proof = ref(null);
const isLoading = ref(true);
const errorMessage = ref("");
const comment = ref("");
const successMessage = ref("");

async function loadProof() {
  try {
    proof.value = await getProofById(route.params.id);
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
    <RouterLink to="/customer/dashboard" class="back-link">
      ← Back to dashboard
    </RouterLink>

    <p v-if="isLoading">Loading proof...</p>

    <section v-else-if="proof" class="review-card">
      <div class="review-header">
        <div>
          <p class="eyebrow">Proof review</p>
          <h1>{{ proof.productName }}</h1>
          <p class="muted">{{ proof.orderId }} · {{ proof.customerName }}</p>
        </div>

        <span class="status-pill">
          {{ proof.status }}
        </span>
      </div>

      <div class="preview-box">
        <p class="preview-title">Proof file preview</p>
        <p class="muted">{{ proof.fileName }}</p>
        <p class="preview-note">
          File preview will be implemented in the upload/proof viewer task.
        </p>
      </div>

      <div class="action-area">
        <button class="approve-button" @click="handleApprove">
          Approve proof
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
          </button>
        </div>
      </div>

      <p v-if="successMessage" class="success">{{ successMessage }}</p>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </section>

    <p v-else class="error">{{ errorMessage }}</p>
  </main>
</template>

<style scoped>
.page {
  min-height: 100vh;
  background: #f7f4f4;
  padding: 32px;
  color: #111827;
  font-family: Inter, Arial, sans-serif;
}

.back-link {
  display: inline-flex;
  margin-bottom: 20px;
  color: #111827;
  font-weight: 800;
  text-decoration: none;
}

.review-card {
  background: white;
  border-radius: 28px;
  padding: 32px;
  max-width: 1000px;
  margin: 0 auto;
}

.review-header {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  margin-bottom: 24px;
}

.eyebrow {
  font-weight: 800;
  margin: 0 0 8px;
}

h1 {
  font-size: 42px;
  margin: 0;
}

.muted {
  color: #6b7280;
}

.status-pill {
  background: #fff3c4;
  color: #8a5b00;
  border-radius: 999px;
  padding: 10px 14px;
  font-weight: 800;
  height: fit-content;
}

.preview-box {
  border: 2px dashed #bdebf4;
  background: #f0fbfd;
  border-radius: 24px;
  padding: 48px;
  text-align: center;
  margin-bottom: 24px;
}

.preview-title {
  font-size: 24px;
  font-weight: 900;
  margin: 0;
}

.preview-note {
  margin-top: 18px;
}

.action-area {
  display: grid;
  grid-template-columns: 240px 1fr;
  gap: 24px;
  align-items: start;
}

.approve-button,
.secondary-button {
  border: none;
  cursor: pointer;
  border-radius: 14px;
  padding: 14px 18px;
  font-weight: 900;
}

.approve-button {
  background: #ff0050;
  color: white;
}

.secondary-button {
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
  font-weight: 900;
}

textarea {
  min-height: 120px;
  border-radius: 16px;
  border: 1px solid #d1d5db;
  padding: 14px;
  font-family: inherit;
}

.success {
  color: #047857;
  font-weight: 800;
}

.error {
  color: #be123c;
  font-weight: 800;
}
</style>
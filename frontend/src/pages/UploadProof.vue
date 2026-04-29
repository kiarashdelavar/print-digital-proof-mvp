<script setup>
import { ref } from "vue";
import { RouterLink } from "vue-router";
import { createProof } from "../services/proofService";
import printLogo from "../assets/print-logo.svg";

const form = ref({
  orderId: "",
  customerName: "",
  customerEmail: "",
  productName: "",
  fileName: "",
});

const createdProof = ref(null);
const isSubmitting = ref(false);
const errorMessage = ref("");

async function handleSubmit() {
  errorMessage.value = "";
  createdProof.value = null;

  if (
    !form.value.orderId ||
    !form.value.customerName ||
    !form.value.customerEmail ||
    !form.value.productName ||
    !form.value.fileName
  ) {
    errorMessage.value = "Please fill in all fields.";
    return;
  }

  try {
    isSubmitting.value = true;
    createdProof.value = await createProof(form.value);

    form.value = {
      orderId: "",
      customerName: "",
      customerEmail: "",
      productName: "",
      fileName: "",
    };
  } catch (error) {
    errorMessage.value = "Could not create proof. Please check if the backend is running.";
  } finally {
    isSubmitting.value = false;
  }
}
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

    <RouterLink to="/file-checker/dashboard" class="back-link">
      ← Back to file checker dashboard
    </RouterLink>

    <section class="upload-shell">
      <aside class="side-panel">
        <p class="eyebrow">File checker workflow</p>
        <h1>
          Upload a<br />
          digital proof
        </h1>
        <p class="side-text">
          Create a proof request for a customer order and generate a review page
          that the customer can approve or reject.
        </p>
      </aside>

      <section class="form-card">
        <p class="section-eyebrow">New proof request</p>
        <h2>Create digital proof</h2>
        <p class="muted">
          This MVP creates a proof record using JSON data. Real file upload will be added later.
        </p>

        <form class="proof-form" @submit.prevent="handleSubmit">
          <label>
            Order ID
            <input v-model="form.orderId" type="text" placeholder="ORD-1003" />
          </label>

          <label>
            Customer name
            <input v-model="form.customerName" type="text" placeholder="Sarah Miller" />
          </label>

          <label>
            Customer email
            <input v-model="form.customerEmail" type="email" placeholder="sarah@example.com" />
          </label>

          <label>
            Product name
            <input v-model="form.productName" type="text" placeholder="Poster A3" />
          </label>

          <label>
            Proof file name
            <input v-model="form.fileName" type="text" placeholder="poster-a3-proof.pdf" />
          </label>

          <button class="primary-button" type="submit" :disabled="isSubmitting">
            {{ isSubmitting ? "Creating proof..." : "Create proof request" }}
            <span>↗</span>
          </button>
        </form>

        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>

        <div v-if="createdProof" class="success-card">
          <p class="section-eyebrow">Proof created</p>
          <h3>{{ createdProof.productName }}</h3>
          <p>
            The proof is now waiting for customer approval.
          </p>

          <RouterLink :to="`/customer/proofs/${createdProof.id}`" class="review-link">
            Open customer review page
          </RouterLink>
        </div>
      </section>
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

.back-link {
  display: flex;
  max-width: 1680px;
  margin: 0 auto 20px;
  color: #071733;
  font-weight: 950;
  text-decoration: none;
}

.upload-shell {
  max-width: 1680px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 0.9fr 1.2fr;
  gap: 28px;
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
  margin: 28px 0 0;
  color: #0c2538;
  font-size: 18px;
  line-height: 1.55;
  font-weight: 650;
}

.form-card {
  background: white;
  border-radius: 28px;
  padding: 42px;
  box-shadow: 0 14px 34px rgba(7, 23, 51, 0.04);
}

.section-eyebrow {
  margin: 0 0 8px;
  color: #ff0050;
  font-weight: 950;
}

h2 {
  margin: 0;
  font-size: 42px;
}

.muted {
  color: #667085;
}

.proof-form {
  margin-top: 28px;
  display: grid;
  gap: 18px;
}

label {
  display: grid;
  gap: 8px;
  font-weight: 950;
}

input {
  border: 1px solid #d1d5db;
  border-radius: 14px;
  padding: 16px;
  font: inherit;
}

input:focus {
  outline: 3px solid rgba(127, 211, 232, 0.35);
  border-color: #7fd3e8;
}

.primary-button {
  border: none;
  cursor: pointer;
  margin-top: 10px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  gap: 12px;
  background: #ff0050;
  color: white;
  border-radius: 14px;
  padding: 16px 22px;
  font-weight: 950;
  box-shadow: 0 12px 24px rgba(255, 0, 80, 0.18);
}

.primary-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.success-card {
  margin-top: 28px;
  border-radius: 22px;
  background: #eaf9fc;
  border: 1px dashed #7fd3e8;
  padding: 24px;
}

.success-card h3 {
  margin: 0;
  font-size: 28px;
}

.review-link {
  display: inline-flex;
  margin-top: 14px;
  color: #ff0050;
  font-weight: 950;
  text-decoration: none;
}

.error {
  color: #be123c;
  font-weight: 900;
  margin-top: 18px;
}

@media (max-width: 1000px) {
  .upload-shell {
    grid-template-columns: 1fr;
  }

  .side-panel {
    min-height: auto;
  }

  h1 {
    font-size: 44px;
  }

  .nav-actions {
    display: none;
  }
}
</style>
import { createRouter, createWebHistory } from "vue-router";
import CustomerDashboard from "../pages/CustomerDashboard.vue";
import ProofReview from "../pages/ProofReview.vue";

const routes = [
  {
    path: "/",
    redirect: "/customer/dashboard",
  },
  {
    path: "/customer/dashboard",
    component: CustomerDashboard,
  },
  {
    path: "/customer/proofs/:id",
    component: ProofReview,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
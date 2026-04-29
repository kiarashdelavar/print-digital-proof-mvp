import { createRouter, createWebHashHistory } from "vue-router";
import CustomerDashboard from "../pages/CustomerDashboard.vue";
import ProofReview from "../pages/ProofReview.vue";
import FileCheckerDashboard from "../pages/FileCheckerDashboard.vue";
import UploadProof from "../pages/UploadProof.vue";
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
  {
  path: "/file-checker/dashboard",
  component: FileCheckerDashboard,
},
{
  path: "/file-checker/upload",
  component: UploadProof,
},
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
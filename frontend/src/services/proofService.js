import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api";

export async function getProofs() {
  const response = await axios.get(`${API_BASE_URL}/proofs`);
  return response.data;
}

export async function getProofById(id) {
  const response = await axios.get(`${API_BASE_URL}/proofs/${id}`);
  return response.data;
}

export async function approveProof(id) {
  const response = await axios.post(`${API_BASE_URL}/proofs/${id}/approve`);
  return response.data;
}

export async function requestChanges(id, comment) {
  const response = await axios.post(`${API_BASE_URL}/proofs/${id}/request-changes`, {
    comment,
  });

  return response.data;
}
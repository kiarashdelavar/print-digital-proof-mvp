# MVP Plan

## Project Title

Digital Proof Approval Workflow MVP

## Problem

Digital proof approval is often handled manually through email. This can make it difficult to track the current status, the latest proof file, and the customer's final approval.

## MVP Scope

The MVP will focus on a simple workflow:

1. A file checker uploads a digital proof.
2. The proof is connected to an order.
3. The customer can review the proof.
4. The customer can approve the proof or request changes.
5. The file checker can see the updated status.

## Roles

### File Checker

- Upload proof
- View proof status
- See customer feedback

### Customer

- View proof
- Approve proof
- Request changes with a comment

## Proof Statuses

- WAITING_FOR_APPROVAL
- APPROVED
- CHANGES_REQUESTED

## Pages

- Home page
- Customer dashboard
- Proof review page
- File checker dashboard
- Upload proof page

## Backend Endpoints

- GET /api/proofs
- GET /api/proofs/:id
- POST /api/proofs
- POST /api/proofs/:id/approve
- POST /api/proofs/:id/request-changes